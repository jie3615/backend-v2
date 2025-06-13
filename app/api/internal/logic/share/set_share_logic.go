package share

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 设置分享
func NewSetShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetShareLogic {
	return &SetShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetShareLogic) SetShare() (resp *types.ShareSetResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
