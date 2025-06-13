package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comereducation"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnbindComerEducationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 解绑用户教育经历
func NewUnbindComerEducationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnbindComerEducationsLogic {
	return &UnbindComerEducationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnbindComerEducationsLogic) UnbindComerEducations(req *types.UnbindComerEducationsRequest) (resp *types.ComerMessageResponse, err error) {
	err = comereducation.DeleteComerEducation(l.svcCtx.DB, uint64(req.ComerEducationId))
	if err != nil {
		return nil, err
	}
	return &types.ComerMessageResponse{
		Message: "success",
	}, nil
}
