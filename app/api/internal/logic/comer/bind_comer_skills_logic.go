package comer

import (
	"context"
	"errors"
	"fmt"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comer"
	"metaLand/data/model/comerskill"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindComerSkillsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 绑定用户技能
func NewBindComerSkillsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindComerSkillsLogic {
	return &BindComerSkillsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindComerSkillsLogic) BindComerSkills(req *types.BindComerSkillsRequest) (resp *types.ComerMessageResponse, err error) {
	// todo: add your logic here and delete this line
	comerInfo, ok := l.ctx.Value("comerInfo").(*comer.Comer)
	if !ok {
		return nil, errors.New("user not found")
	}

	comerSkill := &comerskill.ComerSkill{
		ComerId:     uint64(comerInfo.ID),
		SkillName:   req.SkillName,
		Level:       req.Level,
		Years:       req.Years,
		Description: req.Description,
	}
	id, err := comerskill.CreateComerSkill(l.svcCtx.DB, comerSkill)
	if err != nil {
		return nil, err
	}
	return &types.ComerMessageResponse{
		Message: fmt.Sprintf("success, id: %d", id),
	}, nil
}
