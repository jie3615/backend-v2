package comers

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyComerAddProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 校验用户是否可添加资料
func NewVerifyComerAddProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyComerAddProfileLogic {
	return &VerifyComerAddProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyComerAddProfileLogic) VerifyComerAddProfile() (resp *types.ThirdPartyVerifyResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
