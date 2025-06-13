package comers

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCustomDomainExistenceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询自定义域名是否存在
func NewGetUserCustomDomainExistenceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCustomDomainExistenceLogic {
	return &GetUserCustomDomainExistenceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserCustomDomainExistenceLogic) GetUserCustomDomainExistence(req *types.GetUserCustomDomainExistenceRequest) (resp *types.IsExistResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
