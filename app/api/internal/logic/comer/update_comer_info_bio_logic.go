package comer

import (
	"context"
	"errors"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comer"
	"metaLand/data/model/comerprofile"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateComerInfoBioLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户简介
func NewUpdateComerInfoBioLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateComerInfoBioLogic {
	return &UpdateComerInfoBioLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateComerInfoBioLogic) UpdateComerInfoBio(req *types.UpdateComerInfoBioRequest) (resp *types.ComerMessageResponse, err error) {
	comerInfo, ok := l.ctx.Value("comerInfo").(*comer.Comer)
	if !ok {
		return nil, errors.New("user not found")
	}

	err = comerprofile.UpdateComerProfile(l.svcCtx.DB, uint64(comerInfo.ID), map[string]interface{}{
		"bio": req.Bio,
	})
	if err != nil {
		return nil, err
	}
	return &types.ComerMessageResponse{
		Message: "success",
	}, nil
}
