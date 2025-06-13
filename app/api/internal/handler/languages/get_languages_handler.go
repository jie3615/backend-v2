package languages

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/languages"
	"metaLand/app/api/internal/svc"
)

// 获取语言列表
func GetLanguagesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := languages.NewGetLanguagesLogic(r.Context(), svcCtx)
		resp, err := l.GetLanguages()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
