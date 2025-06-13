package tags

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/tags"
	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
)

// 根据类型获取标签列表
func GetTagsByTagTypeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTagsByTagTypeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tags.NewGetTagsByTagTypeLogic(r.Context(), svcCtx)
		resp, err := l.GetTagsByTagType(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
