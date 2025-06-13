package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model"
	"metaLand/data/model/comerskill"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateComerSkillsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户技能
func NewUpdateComerSkillsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateComerSkillsLogic {
	return &UpdateComerSkillsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateComerSkillsLogic) UpdateComerSkills(req *types.UpdateComerSkillsRequest) (resp *types.ComerMessageResponse, err error) {
	// todo: add your logic here and delete this line
	comerSkill := comerskill.ComerSkill{
		Base: model.Base{
			ID: uint64(req.ComerSkillId),
		},
		ComerId:     uint64(req.ComerId),
		SkillName:   req.SkillName,
		Level:       req.Level,
		Years:       req.Years,
		Description: req.Description,
	}
	err = comerskill.UpdateComerSkill(l.svcCtx.DB, &comerSkill)
	if err != nil {
		return nil, err
	}

	return &types.ComerMessageResponse{
		Message: "success",
	}, nil
}
