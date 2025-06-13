package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comerlanguage"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnbindComerLanguagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 解绑用户语言
func NewUnbindComerLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnbindComerLanguagesLogic {
	return &UnbindComerLanguagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnbindComerLanguagesLogic) UnbindComerLanguages(req *types.UnbindComerLanguagesRequest) (resp *types.ComerMessageResponse, err error) {
	err = comerlanguage.DeleteComerLanguage(l.svcCtx.DB, uint64(req.ComerLanguageId))
	if err != nil {
		return nil, err
	}

	return &types.ComerMessageResponse{
		Message: "success",
	}, nil
}
