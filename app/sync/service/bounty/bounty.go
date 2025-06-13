package bounty

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"                  //以太坊基础库
	ethCommon "github.com/ethereum/go-ethereum/common" //以太坊常用工具
	"github.com/ethereum/go-ethereum/ethclient"        //以太坊客户端
	"github.com/redis/go-redis/v9"                     //redis客户端
	"github.com/zeromicro/go-zero/core/logx"           //日志组件
	"github.com/zeromicro/go-zero/core/threading"      //协程管理
	"math/big"                                         //大整数计算
	"metaLand/app/sync/service/common"
	"metaLand/data/model/bounty"
	chainModel "metaLand/data/model/chain" //链数据模型
	"time"
)

// TaskBounty 赏金任务处理器
type TaskBounty struct {
	ctx *common.ServiceContext //服务上下文
	// 疑问：链上合约数据，是否是一个链只对应一个合约
	info map[uint64]*[]ContractInfo //链上合约数据，key为chainID
}

func NewTaskBounty(ctx *common.ServiceContext) *TaskBounty {
	return &TaskBounty{
		ctx:  ctx,
		info: make(map[uint64]*[]ContractInfo),
	}
}

func (t *TaskBounty) HandlerCreateEvent(chainId uint64, txHash ethCommon.Hash, params []any) {
	ownerAddress := params[0].(ethCommon.Address)
	factoryAddress := params[1].(ethCommon.Address)
	founderAddress := params[2].(ethCommon.Address)
	pars := params[3].(struct {
		DepositToken              ethCommon.Address `json:"deposit_token"`
		DepositTokenIsNative      bool              `json:"deposit_token_is_native"`
		FounderDepositAmount      big.Int           `json:"founder_deposit_amount"`
		ApplicantDepositMinAmount big.Int           `json:"applicant_deposit_min_amount"`
		ApplyDeadline             big.Int           `json:"apply_deadline"`
	})

	logx.Info("create", ownerAddress, factoryAddress, founderAddress, pars)
	// 根据txHash查询bounty数据
	b := bounty.Bounty{}
	err := bounty.GetBountyByTxHashAndChainId(t.ctx.DB, txHash.Hex(), chainId, &b)
	if err != nil {
		logx.Error(err)
		return
	}
	if b.ID == 0 {
		logx.Info("bounty not exists")
		return
	}
	// 更新bounty数据
	// todo 待完善
	b.DepositContract = ownerAddress.Hex()
	err = bounty.UpdateBounty(t.ctx.DB, &b)
	if err != nil {
		logx.Error(err)
		return
	}
}

func (t *TaskBounty) queryLogs() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for chainID, infos := range t.info {
		for _, info := range *infos {
			startBlock := big.NewInt(0)
			bountyABI, err := common.GetABI(info.ABI)
			if err != nil {
				logx.Error(err)
				return
			}
			lastHeigh, err := t.ctx.Redis.Get(ctx, common.GetKey(chainID, info.Address)).Uint64()
			if err != nil && !errors.Is(err, redis.Nil) {
				logx.Error(err)
				continue
			}
			if lastHeigh == 0 {
				receipt, err := info.Client.TransactionReceipt(ctx, ethCommon.HexToHash(info.CreatedHash))
				if err != nil {
					logx.Error(err)
					continue
				}
				startBlock = big.NewInt(0).Add(receipt.BlockNumber, big.NewInt(1))
			} else {
				startBlock = big.NewInt(int64(lastHeigh + 1))
			}
			currentHeight, err := info.Client.BlockNumber(ctx)
			if err != nil {
				logx.Info(err)
				continue
			}
			if big.NewInt(int64(currentHeight)).Cmp(startBlock) <= 0 {
				continue
			}
			endBlock := big.NewInt(0).Add(startBlock, big.NewInt(int64(499)))
			if endBlock.Cmp(big.NewInt(int64(currentHeight))) > 0 {
				endBlock = big.NewInt(int64(currentHeight))
			}

			for {
				logx.Info(fmt.Sprintf("start: %d end: %d current: %d", startBlock.Int64(), endBlock.Int64(), currentHeight))
				logs, err := info.Client.FilterLogs(ctx, ethereum.FilterQuery{
					FromBlock: startBlock,
					ToBlock:   endBlock,
					Addresses: []ethCommon.Address{
						ethCommon.HexToAddress(info.Address),
					},
				})
				if err != nil {
					logx.Info(err)
					break
				}
				for _, l := range logs {
					switch l.Topics[0] {
					case ethCommon.HexToHash(EventCreated):
						params, err := bountyABI.Events["created"].Inputs.UnpackValues(l.Data)
						if err != nil {
							logx.Error(err)
							continue
						}
						t.HandlerCreateEvent(chainID, l.TxHash, params)
						//  todo 待完善(Deposit, Close, Approve, Unapprove, Apply, ReleaseFounderDeposit, ReleaseApplicantDeposit, Lock, Unlock, PostUpdate)
					default:
						logx.Info(l.Topics[0])
					}
				}
				err = t.ctx.Redis.Set(ctx, common.GetKey(chainID, info.Address), endBlock.Uint64(), 0).Err()
				if err != nil {
					logx.Info(err)
					break
				}
				startBlock = big.NewInt(0).Add(endBlock, big.NewInt(int64(1)))
				endBlock = big.NewInt(0).Add(startBlock, big.NewInt(int64(499)))
				if endBlock.Cmp(big.NewInt(int64(currentHeight))) > 0 {
					endBlock = big.NewInt(int64(currentHeight))
				}
				if endBlock.Cmp(startBlock) <= 0 {
					break
				}
			}
		}
	}
}

func (t *TaskBounty) process() {
	chains := make([]chainModel.ChainBasicResponse, 0)
	err := chainModel.GetChainCompleteList(t.ctx.DB, &chains)
	if err != nil {
		logx.Error(err)
	}
	for _, chain := range chains {
		var rpcUrl string
		for _, endpoint := range chain.ChainEndpoints {
			if endpoint.Protocol == 1 {
				rpcUrl = endpoint.URL
			}
		}
		cli, err := ethclient.Dial(rpcUrl)
		if err != nil {
			logx.Error(err)
			continue
		}
		for _, contract := range chain.ChainContracts {
			// bounty contract
			if contract.Project == 4 {
				*t.info[chain.ChainID] = append(*t.info[chain.ChainID], ContractInfo{
					Address:     contract.Address,
					CreatedHash: contract.CreatedTxHash,
					Client:      cli,
					ABI:         contract.ABI,
				})
			}
		}
	}
	for {
		t.queryLogs()
		time.Sleep(3 * time.Second)
	}
}

func (t *TaskBounty) Start() {
	threading.GoSafe(t.process)
}

func (t *TaskBounty) AddBounty(chainId uint64, bountyAddress string, contract ContractInfo) {
	b, err := chainModel.GetChainContractByChainIdAndContractAddress(t.ctx.DB, chainId, bountyAddress)
	if err != nil {
		logx.Error(err)
		return
	}
	*t.info[chainId] = append(*t.info[chainId], ContractInfo{
		Address:     bountyAddress,
		ABI:         b.ABI,
		Client:      contract.Client,
		CreatedHash: b.CreatedTxHash,
	})
}
