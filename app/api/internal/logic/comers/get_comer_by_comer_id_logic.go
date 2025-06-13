package comers

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComerByComerIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 通过comer_id获取用户
func NewGetComerByComerIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetComerByComerIdLogic {
	return &GetComerByComerIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComerByComerIdLogic) GetComerByComerId(req *types.GetComerByComerIdRequest) (resp *types.ComerInfoDetailResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
