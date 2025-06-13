package comers

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComerParticipatedCountByComerIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取该用户参与的项目数量
func NewGetComerParticipatedCountByComerIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetComerParticipatedCountByComerIdLogic {
	return &GetComerParticipatedCountByComerIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComerParticipatedCountByComerIdLogic) GetComerParticipatedCountByComerId(req *types.GetComerParticipatedCountByComerIdRequest) (resp *types.ProjectCountResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
