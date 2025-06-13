package comer

import (
	"context"
	"errors"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comer"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComerInvitationRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户邀请记录
func NewGetComerInvitationRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetComerInvitationRecordsLogic {
	return &GetComerInvitationRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComerInvitationRecordsLogic) GetComerInvitationRecords() (resp *types.ComerPageData, err error) {
	comerInfo, ok := l.ctx.Value("comerInfo").(*comer.Comer)
	if !ok {
		return nil, errors.New("user not found")
	}

	logx.Infof("comerInfo: %+v", comerInfo)

	return &types.ComerPageData{
		Page:  1,
		Size:  10,
		Total: 0,
		List:  []interface{}{},
	}, nil
}
