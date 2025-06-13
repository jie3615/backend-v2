package comer

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/comer"
	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
)

// 绑定用户社交
func BindComerSocialsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BindComerSocialsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comer.NewBindComerSocialsLogic(r.Context(), svcCtx)
		resp, err := l.BindComerSocials(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
