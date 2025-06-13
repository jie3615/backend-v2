package comer

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/comer"
	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
)

// 解绑用户账号
func UnlinkOauthByComerAccountIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UnlinkOauthByComerAccountIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comer.NewUnlinkOauthByComerAccountIdLogic(r.Context(), svcCtx)
		resp, err := l.UnlinkOauthByComerAccountId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
