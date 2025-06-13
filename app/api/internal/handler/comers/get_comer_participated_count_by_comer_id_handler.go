package comers

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/comers"
	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
)

// 获取该用户参与的项目数量
func GetComerParticipatedCountByComerIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetComerParticipatedCountByComerIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comers.NewGetComerParticipatedCountByComerIdLogic(r.Context(), svcCtx)
		resp, err := l.GetComerParticipatedCountByComerId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
