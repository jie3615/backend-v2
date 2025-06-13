package startup

import (
	"context"
	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/startup"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStartupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建项目

func NewCreateStartupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStartupsLogic {
	return &CreateStartupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStartupsLogic) CreateStartups(req *types.CreateStartupsRequest) (resp *types.CreateStartupsResponse, err error) {

	var response types.CreateStartupsResponse
	// 创建完整的请求对象，映射所有字段
	createReq := &startup.CreateStartupsRequest{
		ComerID:              req.ComerID,
		Name:                 req.Name,
		Mode:                 req.Mode,
		Logo:                 req.Logo,
		Cover:                req.Cover,
		Mission:              req.Mission,
		TokenContractAddress: req.TokenContractAddress,
		Overview:             req.Overview,
		TxHash:               req.TxHash,
		OnChain:              req.OnChain,
		KYC:                  req.KYC,
		ContractAudit:        req.ContractAudit,
		Website:              req.Website,
		Discord:              req.Discord,
		Twitter:              req.Twitter,
		Telegram:             req.Telegram,
		Docs:                 req.Docs,
		Email:                req.Email,
		Facebook:             req.Facebook,
		Medium:               req.Medium,
		Linktree:             req.Linktree,
		LaunchNetwork:        req.LaunchNetwork,
		TokenName:            req.TokenName,
		TokenSymbol:          req.TokenSymbol,
		TotalSupply:          req.TotalSupply,
		PresaleStart:         req.PresaleStart,
		PresaleEnd:           req.PresaleEnd,
		LaunchDate:           req.LaunchDate,
		TabSequence:          req.TabSequence,
		IsDeleted:            req.IsDeleted,
	}

	// 调用数据层创建方法
	data, err := startup.CreateStartups(l.svcCtx.DB, createReq)

	if err != nil {
		return nil, err
	}

	response.Suc = data
	response.Msg = "写入成功!"

	return &types.CreateStartupsResponse{
		Suc: true,
		Msg: "写入成功",
	}, nil
}
