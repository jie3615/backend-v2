package comers

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComerPostedCountByComerIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取该用户发布的项目数量
func NewGetComerPostedCountByComerIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetComerPostedCountByComerIdLogic {
	return &GetComerPostedCountByComerIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComerPostedCountByComerIdLogic) GetComerPostedCountByComerId(req *types.GetComerPostedCountByComerIdRequest) (resp *types.ProjectCountResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
