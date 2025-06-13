package comers

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComerConnectComersByComerIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取该用户连接的comer列表
func NewGetComerConnectComersByComerIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetComerConnectComersByComerIdLogic {
	return &GetComerConnectComersByComerIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComerConnectComersByComerIdLogic) GetComerConnectComersByComerId(req *types.GetComerConnectComersByComerIdRequest) (resp *types.ComersPageData, err error) {
	// todo: add your logic here and delete this line

	return
}
