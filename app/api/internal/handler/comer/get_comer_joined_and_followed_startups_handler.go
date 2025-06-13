package comer

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/comer"
	"metaLand/app/api/internal/svc"
)

// 获取用户加入和关注的创业公司
func GetComerJoinedAndFollowedStartupsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := comer.NewGetComerJoinedAndFollowedStartupsLogic(r.Context(), svcCtx)
		resp, err := l.GetComerJoinedAndFollowedStartups()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
