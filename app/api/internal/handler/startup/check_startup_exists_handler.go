package startup

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/startup"
	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
)

// 判断项目是否存在
func CheckStartupExistsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckStartupExistsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := startup.NewCheckStartupExistsLogic(r.Context(), svcCtx)
		resp, err := l.CheckStartupExists(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
