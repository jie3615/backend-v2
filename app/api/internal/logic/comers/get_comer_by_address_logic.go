package comers

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComerByAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 通过地址获取用户
func NewGetComerByAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetComerByAddressLogic {
	return &GetComerByAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComerByAddressLogic) GetComerByAddress(req *types.GetComerByAddressRequest) (resp *types.ComerBasicResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
