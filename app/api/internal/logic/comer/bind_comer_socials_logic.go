package comer

import (
	"context"
	"errors"
	"fmt"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comer"
	"metaLand/data/model/comersocial"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindComerSocialsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 绑定用户社交
func NewBindComerSocialsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindComerSocialsLogic {
	return &BindComerSocialsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindComerSocialsLogic) BindComerSocials(req *types.BindComerSocialsRequest) (resp *types.ComerMessageResponse, err error) {
	comerInfo, ok := l.ctx.Value("comerInfo").(*comer.Comer)
	if !ok {
		return nil, errors.New("user not found")
	}

	comerSocial := &comersocial.ComerSocial{
		ComerId:    uint64(comerInfo.ID),
		Platform:   req.PlatformName,
		Username:   req.UserName,
		Url:        req.Url,
		IsVerified: req.IsVerified,
	}

	id, err := comersocial.CreateComerSocial(l.svcCtx.DB, comerSocial)
	if err != nil {
		return nil, err
	}

	return &types.ComerMessageResponse{
		Message: fmt.Sprintf("success, id: %d", id),
	}, nil
}
