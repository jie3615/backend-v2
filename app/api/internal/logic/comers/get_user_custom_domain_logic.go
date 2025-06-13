package comers

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCustomDomainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 通过自定义域名获取用户
func NewGetUserCustomDomainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCustomDomainLogic {
	return &GetUserCustomDomainLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserCustomDomainLogic) GetUserCustomDomain(req *types.GetUserCustomDomainRequest) (resp *types.ComerInfoDetailResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
