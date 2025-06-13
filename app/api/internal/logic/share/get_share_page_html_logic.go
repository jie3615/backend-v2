package share

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"metaLand/app/api/internal/svc"
)

type GetSharePageHtmlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取分享
func NewGetSharePageHtmlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSharePageHtmlLogic {
	return &GetSharePageHtmlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSharePageHtmlLogic) GetSharePageHtml() (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
