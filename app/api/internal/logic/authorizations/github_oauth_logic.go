package authorizations

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GithubOauthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Github 授权登录
func NewGithubOauthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GithubOauthLogic {
	return &GithubOauthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GithubOauthLogic) GithubOauth() (resp *types.JwtAuthorizationResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
