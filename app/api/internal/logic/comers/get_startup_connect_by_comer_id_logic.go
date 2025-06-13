package comers

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStartupConnectByComerIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取该用户连接的startup列表
func NewGetStartupConnectByComerIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStartupConnectByComerIdLogic {
	return &GetStartupConnectByComerIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStartupConnectByComerIdLogic) GetStartupConnectByComerId(req *types.GetComerConnectStartupsByComerIdRequest) (resp *types.ComersPageData, err error) {
	// todo: add your logic here and delete this line

	return
}
