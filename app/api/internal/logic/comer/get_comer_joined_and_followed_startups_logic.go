package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComerJoinedAndFollowedStartupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户加入和关注的创业公司
func NewGetComerJoinedAndFollowedStartupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetComerJoinedAndFollowedStartupsLogic {
	return &GetComerJoinedAndFollowedStartupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComerJoinedAndFollowedStartupsLogic) GetComerJoinedAndFollowedStartups() (resp *types.StartupListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
