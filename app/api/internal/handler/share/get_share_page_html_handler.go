package share

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/share"
	"metaLand/app/api/internal/svc"
)

// 获取分享
func GetSharePageHtmlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := share.NewGetSharePageHtmlLogic(r.Context(), svcCtx)
		resp, err := l.GetSharePageHtml()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
