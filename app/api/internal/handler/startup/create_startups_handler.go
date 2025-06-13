package startup

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/startup"
	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
)

// 创建项目
func CreateStartupsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateStartupsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := startup.NewCreateStartupsLogic(r.Context(), svcCtx)
		resp, err := l.CreateStartups(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
