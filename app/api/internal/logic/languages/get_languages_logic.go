package languages

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLanguagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取语言列表
func NewGetLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLanguagesLogic {
	return &GetLanguagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLanguagesLogic) GetLanguages() (resp *types.LanguageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
