package authorizations

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/authorizations"
	"metaLand/app/api/internal/svc"
)

// Google 授权登录
func GoogleOauthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := authorizations.NewGoogleOauthLogic(r.Context(), svcCtx)
		resp, err := l.GoogleOauth()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
