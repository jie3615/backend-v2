package comers

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComerBeConnectComersByComerIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取被该用户连接的comer列表
func NewGetComerBeConnectComersByComerIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetComerBeConnectComersByComerIdLogic {
	return &GetComerBeConnectComersByComerIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComerBeConnectComersByComerIdLogic) GetComerBeConnectComersByComerId(req *types.GetComerBeConnectComersByComerIdRequest) (resp *types.ComersPageData, err error) {
	// todo: add your logic here and delete this line

	return
}
