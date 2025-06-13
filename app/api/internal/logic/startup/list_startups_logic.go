package startup

import (
	"context"
	"metaLand/data/model/startup"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListStartupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询项目列表

func NewListStartupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListStartupsLogic {
	return &ListStartupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListStartupsLogic) ListStartups(req *types.ListStartupsRequest) (resp *types.ListStartupsResponse, err error) {

	var response types.ListStartupsResponse
	startups := make([]startup.Startup, 0)

	//请求参数
	total, err := startup.ListStartups(l.svcCtx.DB, 0, &startup.ListStartupRequest{
		Limit:     req.Limit,
		Offset:    req.Offset,
		IsDeleted: req.IsDeleted,
		Keyword:   req.Keyword,
		Mode:      req.Mode,
	}, &startups)

	if err != nil {
		return nil, err
	}

	response.Total = total

	for _, s := range startups {
		response.List = append(response.List, &types.Startup{
			ComerID:              s.ComerID,
			Name:                 s.Name,
			Mode:                 s.Mode,
			Logo:                 s.Logo,
			Cover:                s.Cover,
			Mission:              s.Mission,
			TokenContractAddress: s.TokenContractAddress,
			Overview:             s.Overview,
			//ChainID:              s.ChainID,
			TxHash:        s.TxHash,
			OnChain:       s.OnChain,
			KYC:           s.KYC,
			ContractAudit: s.ContractAudit,
		})
	}
	return
}
