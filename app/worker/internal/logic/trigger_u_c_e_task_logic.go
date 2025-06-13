package logic

import (
	"context"

	"metaLand/app/worker/internal/svc"
	"metaLand/app/worker/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type TriggerUCETaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTriggerUCETaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TriggerUCETaskLogic {
	return &TriggerUCETaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TriggerUCETaskLogic) TriggerUCETask(in *pb.TriggerUCETaskRequest) (*pb.TriggerUCETaskResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.TriggerUCETaskResponse{}, nil
}
