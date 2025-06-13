package middleware

import (
	"context"
	"metaLand/app/api/internal/config"
	"metaLand/app/api/utility/jwt"
	"metaLand/data/model/comer"
	"net/http"

	"gorm.io/gorm"
)

type OIDCAuthMiddleware struct {
	config config.Config
	DB     *gorm.DB
}

const (
	AuthorizationHeader = "X-COMUNION-AUTHORIZATION"
	ComerUinContextKey  = "COMUNIONCOMERUIN"
	ComerRoleContextKey = "COMUNIONROLE"
	ComerGuestRole      = "Guest"
	ComerLoginedRole    = "Comer"
)

func NewOIDCAuthMiddleware(config config.Config, db *gorm.DB) *OIDCAuthMiddleware {
	return &OIDCAuthMiddleware{
		config: config,
		DB:     db,
	}
}

func (m *OIDCAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		token := r.Header.Get(AuthorizationHeader)
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		comerID, err := jwt.Verify(token, m.config.JWT.Secret)
		if err != nil {
			http.Error(w, "Verify token failed", http.StatusUnauthorized)
			return
		}

		comer, err := comer.FindComer(m.DB, uint64(comerID))
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		// 传递用户信息到上下文
		ctx := context.WithValue(r.Context(), ComerUinContextKey, comerID)
		ctx = context.WithValue(ctx, ComerRoleContextKey, ComerLoginedRole)
		ctx = context.WithValue(ctx, "comerInfo", comer)
		// Passthrough to next handler if need
		next(w, r.WithContext(ctx))
	}
}
