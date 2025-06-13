package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comersocial"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnbindComerSocialsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 解绑用户社交
func NewUnbindComerSocialsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnbindComerSocialsLogic {
	return &UnbindComerSocialsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnbindComerSocialsLogic) UnbindComerSocials(req *types.UnbindComerSocialsRequest) (resp *types.ComerMessageResponse, err error) {
	err = comersocial.DeleteComerSocial(l.svcCtx.DB, uint64(req.ComerSocialId))
	if err != nil {
		return nil, err
	}
	return &types.ComerMessageResponse{
		Message: "success",
	}, nil
}
