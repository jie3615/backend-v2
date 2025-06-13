package chain

import (
	"context"
	"encoding/json"
	"metaLand/data/model/chain"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChainListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取链列表
func NewGetChainListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChainListLogic {
	return &GetChainListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChainListLogic) GetChainList() (resp *types.ChainListResponse, err error) {
	var list chain.ChainListResponse
	err = chain.GetChainCompleteList(l.svcCtx.DB, &list.List)
	if err != nil {
		return
	}

	data, err := json.Marshal(list)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}
