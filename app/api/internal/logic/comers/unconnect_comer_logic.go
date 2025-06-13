package comers

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnconnectComerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 取消连接某个comer
func NewUnconnectComerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnconnectComerLogic {
	return &UnconnectComerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnconnectComerLogic) UnconnectComer(req *types.UnconnectComerRequest) (resp *types.ComersMessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
