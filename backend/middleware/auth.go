package middleware

import (
	"backend/ctx"
	"backend/pkg/jwt"
	"net/http"
	"strings"
)

type Middleware func(next http.Handler) http.Handler

const (
	bearer = "Bearer "
)

// TODO: check jwt.JWT external dependency
func Auth(j *jwt.JWT) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			token = strings.ReplaceAll(token, bearer, "")
			claim, err := j.Validate(token)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return

			}
			ctx := ctx.WithUsername(r.Context(), claim.Username)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
