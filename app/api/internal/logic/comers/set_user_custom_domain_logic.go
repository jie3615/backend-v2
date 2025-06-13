package comers

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserCustomDomainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 设置用户自定义域名
func NewSetUserCustomDomainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserCustomDomainLogic {
	return &SetUserCustomDomainLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserCustomDomainLogic) SetUserCustomDomain(req *types.SetUserCustomDomainRequest) (resp *types.ComersMessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
