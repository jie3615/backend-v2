package comer

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/comer"
	"metaLand/app/api/internal/svc"
)

// 获取用户邀请人数
func GetComerInvitationCountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := comer.NewGetComerInvitationCountLogic(r.Context(), svcCtx)
		resp, err := l.GetComerInvitationCount()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
