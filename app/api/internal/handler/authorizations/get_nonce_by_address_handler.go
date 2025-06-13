package authorizations

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"metaLand/app/api/internal/logic/authorizations"
	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
)

// 获取钱包地址登录用的 nonce
func GetNonceByAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetNonceByAddressRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := authorizations.NewGetNonceByAddressLogic(r.Context(), svcCtx)
		resp, err := l.GetNonceByAddress(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
