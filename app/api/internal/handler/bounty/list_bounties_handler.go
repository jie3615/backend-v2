package bounty

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/bounty"
	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
)

// 查询bounty列表
func ListBountiesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListBountiesRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := bounty.NewListBountiesLogic(r.Context(), svcCtx)
		resp, err := l.ListBounties(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
