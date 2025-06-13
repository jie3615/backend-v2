package comer

import (
	"context"
	"errors"
	"fmt"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comer"
	"metaLand/data/model/comerlanguage"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindComerLanguagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 绑定用户语言
func NewBindComerLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindComerLanguagesLogic {
	return &BindComerLanguagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindComerLanguagesLogic) BindComerLanguages(req *types.BindComerLanguagesRequest) (resp *types.ComerMessageResponse, err error) {
	comerInfo, ok := l.ctx.Value("comerInfo").(*comer.Comer)
	if !ok {
		return nil, errors.New("user not found")
	}

	comerLanguage := &comerlanguage.ComerLanguage{
		ComerId:  uint64(comerInfo.ID),
		Language: req.Language,
		Code:     req.Code,
		Level:    req.Level,
		IsNative: req.IsNative,
	}

	id, err := comerlanguage.CreateComerLanguage(l.svcCtx.DB, comerLanguage)
	if err != nil {
		return nil, err
	}

	return &types.ComerMessageResponse{
		Message: fmt.Sprintf("success, id: %d", id),
	}, nil

}
