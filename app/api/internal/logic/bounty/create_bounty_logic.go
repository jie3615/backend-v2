package bounty

import (
	"context"
	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/bounty"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBountyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建bounty
func NewCreateBountyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBountyLogic {
	return &CreateBountyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateBountyLogic) CreateBounty(req *types.CreateBountyRequest) (resp *types.CreateBountyResponse, err error) {

	//  添加bounty，内容待完善
	err = bounty.AddBounty(l.svcCtx.DB, &bounty.Bounty{
		Title:   req.Title,
		TxHash:  req.TxHash,
		ChainId: req.ChainId,
	})
	if err != nil {
		return
	}
	return &types.CreateBountyResponse{
		Msg:     "success",
		Success: true,
	}, nil
}
