package startup

import (
	"context"
	"fmt"
	_ "metaLand/data/model"
	"metaLand/data/model/startup"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStartupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新项目

func NewUpdateStartupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStartupsLogic {
	return &UpdateStartupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStartupsLogic) UpdateStartups(req *types.UpdateStartupsRequest) (resp *types.UpdateStartupsResponse, err error) {

	// 检查 startupId 是否为空指针
	if req.StartupId < 0 {
		return nil, fmt.Errorf("startupId is required")
	}

	// 参数校验
	if req.Name == "" {
		return &types.UpdateStartupsResponse{
			Suc: false,
			Msg: "项目名称不能为空",
		}, nil
	}

	if req.Overview == "" {
		return &types.UpdateStartupsResponse{
			Suc: false,
			Msg: "项目简介不能为空",
		}, nil
	}
	startupId := req.StartupId

	// 查询项目详情
	startupInfo, err := startup.GetStartupInfo(l.svcCtx.DB, &startupId)
	if err != nil {
		return nil, fmt.Errorf("查询项目失败: %w", err)
	}

	// 更新项目信息
	if req.Name != "" {
		startupInfo.Name = req.Name
	}
	if req.Mode != 0 {
		startupInfo.Mode = req.Mode
	}
	if req.Logo != "" {
		startupInfo.Logo = req.Logo
	}
	if req.Cover != "" {
		startupInfo.Cover = req.Cover
	}
	if req.Overview != "" {
		startupInfo.Overview = req.Overview
	}
	if req.Website != "" {
		startupInfo.Website = req.Website
	}

	// 执行更新
	data, err := startup.UpdateStartup(l.svcCtx.DB, startupInfo)

	if err != nil {
		return &types.UpdateStartupsResponse{
			Suc: false,
			Msg: "项目更新失败",
		}, err
	}

	return &types.UpdateStartupsResponse{
		Suc: data,
		Msg: "项目更新成功",
	}, nil
}
