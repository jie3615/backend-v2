package comers

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConnectedComerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取该用户的连接状态
func NewConnectedComerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConnectedComerLogic {
	return &ConnectedComerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConnectedComerLogic) ConnectedComer(req *types.GetConnectedComerRequest) (resp *types.IsConnectedResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
