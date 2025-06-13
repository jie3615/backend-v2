package tags

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagsByTagTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据类型获取标签列表
func NewGetTagsByTagTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagsByTagTypeLogic {
	return &GetTagsByTagTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTagsByTagTypeLogic) GetTagsByTagType(req *types.GetTagsByTagTypeRequest) (resp *types.ListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
