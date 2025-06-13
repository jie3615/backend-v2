package comer

import (
	"context"
	"errors"
	"fmt"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comer"
	"metaLand/data/model/comereducation"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindComerEducationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 绑定用户教育经历
func NewBindComerEducationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindComerEducationsLogic {
	return &BindComerEducationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindComerEducationsLogic) BindComerEducations(req *types.BindComerEducationsRequest) (resp *types.ComerMessageResponse, err error) {
	comerInfo, ok := l.ctx.Value("comerInfo").(*comer.Comer)
	if !ok {
		return nil, errors.New("user not found")
	}

	comerEducation := &comereducation.ComerEducation{
		ComerId:     uint64(comerInfo.ID),
		School:      req.School,
		Major:       req.Major,
		Degree:      req.Degree,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Description: req.Description,
	}

	id, err := comereducation.CreateComerEducation(l.svcCtx.DB, comerEducation)
	if err != nil {
		return nil, err
	}
	return &types.ComerMessageResponse{
		Message: fmt.Sprintf("%d", id),
	}, nil
}
