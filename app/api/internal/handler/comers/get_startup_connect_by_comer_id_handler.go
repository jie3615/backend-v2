package comers

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/comers"
	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
)

// 获取该用户连接的startup列表
func GetStartupConnectByComerIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetComerConnectStartupsByComerIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comers.NewGetStartupConnectByComerIdLogic(r.Context(), svcCtx)
		resp, err := l.GetStartupConnectByComerId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
