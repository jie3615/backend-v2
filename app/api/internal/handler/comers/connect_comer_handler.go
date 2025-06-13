package comers

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/comers"
	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
)

// 连接某个comer
func ConnectComerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ConnectComerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comers.NewConnectComerLogic(r.Context(), svcCtx)
		resp, err := l.ConnectComer(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
