package comers

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComerInfoDetailByComerIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 通过comer_id获取用户详情
func NewGetComerInfoDetailByComerIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetComerInfoDetailByComerIdLogic {
	return &GetComerInfoDetailByComerIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComerInfoDetailByComerIdLogic) GetComerInfoDetailByComerId(req *types.GetComerInfoDetailByComerIdRequest) (resp *types.ComerInfoDetailResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
