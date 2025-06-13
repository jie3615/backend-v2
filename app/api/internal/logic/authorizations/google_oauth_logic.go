package authorizations

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoogleOauthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Google 授权登录
func NewGoogleOauthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoogleOauthLogic {
	return &GoogleOauthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoogleOauthLogic) GoogleOauth() (resp *types.JwtAuthorizationResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
