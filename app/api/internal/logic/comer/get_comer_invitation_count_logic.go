package comer

import (
	"context"
	"errors"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comer"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComerInvitationCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户邀请人数
func NewGetComerInvitationCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetComerInvitationCountLogic {
	return &GetComerInvitationCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComerInvitationCountLogic) GetComerInvitationCount() (resp *types.ComerInvitationCountResponse, err error) {
	comerInfo, ok := l.ctx.Value("comerInfo").(*comer.Comer)
	if !ok {
		return nil, errors.New("user not found")
	}
	logx.Infof("comerInfo: %+v", comerInfo)
	return &types.ComerInvitationCountResponse{
		ActivatedTotal: 0,
		InactiveTotal:  0,
	}, nil
}
