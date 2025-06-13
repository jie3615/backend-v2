package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comerskill"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnbindComerSkillsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 解绑用户技能
func NewUnbindComerSkillsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnbindComerSkillsLogic {
	return &UnbindComerSkillsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnbindComerSkillsLogic) UnbindComerSkills(req *types.UnbindComerSkillsRequest) (resp *types.ComerMessageResponse, err error) {
	err = comerskill.DeleteComerSkill(l.svcCtx.DB, uint64(req.ComerSkillId))
	if err != nil {
		return nil, err
	}
	return &types.ComerMessageResponse{
		Message: "success",
	}, nil
}
