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
	chainModel "metaLand/data/model/chain" //链数据模型
	"time"
)

// TaskBounty 赏金任务处理器
type TaskBountyFactory struct {
	ctx *common.ServiceContext //服务上下文
	// 疑问：链上合约数据，是否是一个链只对应一个合约
	info       map[uint64]*ContractInfo //链上合约数据，key为chainID
	TaskBounty TaskBounty
}

func NewTaskBountyFactory(ctx *common.ServiceContext) *TaskBountyFactory {
	return &TaskBountyFactory{
		ctx:  ctx,
		info: make(map[uint64]*ContractInfo),
	}
}

func (t *TaskBountyFactory) HandlerCreateEvent(chainId uint64, txHash ethCommon.Hash, params []any) {
	bountyAddress := params[1].(string)
	logx.Info("bountyFactory create", bountyAddress)
	// 入库
	t.ctx.DB.Create(&chainModel.ChainContract{
		ChainID:       chainId,
		Address:       bountyAddress,
		CreatedTxHash: txHash.Hex(),
		Project:       4,
		Type:          1,
		Version:       "",
		ABI:           "", // 疑问：abi文件如何获取
	})
	// 向taskbounty中添加数据
	t.TaskBounty.AddBounty(chainId, bountyAddress, *t.info[chainId])
}

func (t *TaskBountyFactory) queryLogs() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for chainID, info := range t.info {
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

func (t *TaskBountyFactory) process() {
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
			// bountyFactory contract
			if contract.Project == 2 {
				_, has := t.info[chain.ChainID]
				if !has {
					t.info[chain.ChainID] = &ContractInfo{
						Address:     contract.Address,
						CreatedHash: contract.CreatedTxHash,
						Client:      cli,
						ABI:         contract.ABI,
					}
				}
			}
		}
	}
	for {
		t.queryLogs()
		time.Sleep(3 * time.Second)
	}
}

func (t *TaskBountyFactory) Start() {
	threading.GoSafe(t.process)
}
