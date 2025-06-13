package startup

import (
	"context"
	"metaLand/data/model/startup"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckStartupExistsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 判断项目是否存在

func NewCheckStartupExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckStartupExistsLogic {
	return &CheckStartupExistsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckStartupExistsLogic) CheckStartupExists(req *types.CheckStartupExistsRequest) (resp *types.CheckStartupExistsResponse, err error) {
	// 校验入参：至少传一个唯一标识字段
	if req.Name == "" && req.TokenContractAddress == "" {
		return &types.CheckStartupExistsResponse{Exists: false}, nil
	}
	checkReq := &startup.CheckStartupsRequest{
		TokenContractAddress: req.TokenContractAddress,
		Name:                 req.Name,
		IsDeleted:            req.IsDeleted,
	}
	// 构建查询条件
	data, err := startup.CheckExists(l.svcCtx.DB, checkReq)
	if err != nil {
		return nil, err
	}

	return &types.CheckStartupExistsResponse{Exists: data}, nil
}
