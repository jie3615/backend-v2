package bounty

import (
	"context"
	"metaLand/data/model/bounty"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBountiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询bounty列表
func NewListBountiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBountiesLogic {
	return &ListBountiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListBountiesLogic) ListBounties(req *types.ListBountiesRequest) (resp *types.ListBountiesResponse, err error) {
	var response types.ListBountiesResponse
	bounties := make([]bounty.Bounty, 0)
	total, err := bounty.ListBounties(l.svcCtx.DB, &bounty.ListBountiesRequest{
		IsDeleted: req.IsDeleted,
		Keyword:   req.Keyword,
		Limit:     req.Limit,
		Offset:    req.Offset,
	}, &bounties)
	if err != nil {
		return nil, err
	}
	response.Total = total
	for _, v := range bounties {
		response.List = append(response.List, &types.Bounty{
			ApplicantDeposit:            v.ApplicantDeposit,
			ChainId:                     v.ChainId,
			ComerId:                     v.ComerId,
			CreatedAt:                   v.CreatedAt.Format("2006-01-02"),
			DepositContractAddress:      v.DepositContract, //合约地址
			DepositContractTokenDecimal: v.FounderDeposit,
			DepositContractTokenSymbol:  v.DepositTokenSymbol, //质押代币符号
			DiscussionLink:              v.DiscussionLink,     //讨论链接
			//ExpiredTime: v.
			FounderDeposit: v.FounderDeposit, //创始人质押金额
			Id:             v.ID,
			//IsLock:        v.IsDeleted,
			PaymentMode: v.PaymentMode, //支付方式
			Reward: types.BountyReward{
				BountyId:     v.BountyPaymentPeriod.BountyId,
				Token1Symbol: v.BountyPaymentPeriod.Token1Symbol,
				Token2Symbol: v.BountyPaymentPeriod.Token2Symbol,
				Token1Amount: v.BountyPaymentPeriod.Token1Amount,
				Token2Amount: v.BountyPaymentPeriod.Token2Amount,
			},
			//Skills:         tagRelationResponses,
			Startup: types.StartupBasic{
				Id:   v.Startup.ID,
				Name: v.Startup.Name,
				Logo: v.Startup.Logo,
			},
			StartupId: v.StartupId,
			Status:    v.Status, //任务状态
			Title:     v.Title,
			TxHash:    v.TxHash, //交易hash
		})
	}
	return &response, nil
}
