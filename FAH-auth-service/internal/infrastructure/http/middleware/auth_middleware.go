package middleware

import (
	"context"
	"net/http"

	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/cookies"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/jwt"
)

type MiddlewareService struct {
	jwt    *jwt.JWTserv
	cookie *cookies.CookiesService
}
type contextKey string

const JWTClaimsKey contextKey = "claims"

func NewMiddlewareService(jwt *jwt.JWTserv, cookie *cookies.CookiesService) *MiddlewareService {
	return &MiddlewareService{
		jwt:    jwt,
		cookie: cookie,
	}
}
func (m *MiddlewareService) AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token, err := m.cookie.ExtractTokenFromCookie(r, "accessToken")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		claims, err := m.jwt.ValidateAccessToken(r.Context(), token)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), JWTClaimsKey, claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
func (m *MiddlewareService) RequireRoles(roles []int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		fn := func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value("claims").(*jwt.Claims)
			if !ok || claims == nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			for _, role := range roles {
				if claims.Role == role {
					next.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
		return http.HandlerFunc(fn)
	}
}
