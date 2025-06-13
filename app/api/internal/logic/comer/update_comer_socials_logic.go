package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model"
	"metaLand/data/model/comersocial"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateComerSocialsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户社交
func NewUpdateComerSocialsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateComerSocialsLogic {
	return &UpdateComerSocialsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateComerSocialsLogic) UpdateComerSocials(req *types.UpdateComerSocialsRequest) (resp *types.ComerMessageResponse, err error) {
	comerSocial := comersocial.ComerSocial{
		Base: model.Base{
			ID: uint64(req.ComerSocialId),
		},
		ComerId:    uint64(req.ComerId),
		Platform:   req.PlatformName,
		Username:   req.UserName,
		Url:        req.Url,
		IsVerified: req.IsVerified,
	}
	err = comersocial.UpdateComerSocial(l.svcCtx.DB, &comerSocial)
	if err != nil {
		return nil, err
	}
	return &types.ComerMessageResponse{
		Message: "success",
	}, nil
}
