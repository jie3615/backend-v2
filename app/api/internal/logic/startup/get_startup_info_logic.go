package startup

import (
	"context"
	"errors"
	"metaLand/data/model/startup"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStartupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取项目详情

func NewGetStartupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStartupInfoLogic {
	return &GetStartupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStartupInfoLogic) GetStartupInfo(req *types.GetStartupInfoRequest) (resp *types.StartupInfoResponse, err error) {

	startupId := req.StartupId
	if startupId < 1 {
		return nil, errors.New("invalid startupId")
	}

	// 查询项目详情
	startupInfo, err := startup.GetStartupInfo(l.svcCtx.DB, &startupId)
	if err != nil {
		return nil, err
	}
	data := &types.Startup{
		ComerID:              startupInfo.ComerID,
		Name:                 startupInfo.Name,
		Mode:                 startupInfo.Mode,
		Logo:                 startupInfo.Logo,
		Cover:                startupInfo.Cover,
		Mission:              startupInfo.Mission,
		TokenContractAddress: startupInfo.TokenContractAddress,
		Overview:             startupInfo.Overview,
		TxHash:               startupInfo.TxHash,
		OnChain:              startupInfo.OnChain,
		KYC:                  startupInfo.KYC,
		ContractAudit:        startupInfo.ContractAudit,
	}
	// 返回成功响应
	return &types.StartupInfoResponse{
		Code:    200,
		Message: "查询成功",
		Data:    *data,
	}, nil
}
