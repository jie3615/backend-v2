package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model"
	"metaLand/data/model/comerlanguage"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateComerLanguagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户语言
func NewUpdateComerLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateComerLanguagesLogic {
	return &UpdateComerLanguagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateComerLanguagesLogic) UpdateComerLanguages(req *types.UpdateComerLanguagesRequest) (resp *types.ComerMessageResponse, err error) {
	comerLanguage := comerlanguage.ComerLanguage{
		Base: model.Base{
			ID: uint64(req.ComerLanguageId),
		},
		ComerId:  uint64(req.ComerId),
		Language: req.Language,
		Code:     req.Code,
		Level:    req.Level,
		IsNative: req.IsNative,
	}
	err = comerlanguage.UpdateComerLanguage(l.svcCtx.DB, &comerLanguage)
	if err != nil {
		return nil, err
	}

	return &types.ComerMessageResponse{
		Message: "success",
	}, nil
}
