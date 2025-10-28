package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gowalillah/ecommerce/util"
)

func (m *Middlewares) AuthenticateJwt(requiredRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			header := r.Header.Get("Authorization")
			if header == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			headerArr := strings.Split(header, " ")
			if len(headerArr) != 2 {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			jwtToken := headerArr[1]

			isVerify, err := util.Verify(jwtToken, m.cnf.JwtSecretKey)
			if err != nil || !isVerify {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			usr, err := util.DecodeToken(jwtToken, m.cnf.JwtSecretKey)
			if err != nil {
				http.Error(w, "Failed to verify token", http.StatusInternalServerError)
				return
			}

			// ✅ Check role-based access
			if len(requiredRoles) > 0 {
				allowed := false
				for _, role := range requiredRoles {
					if usr.Role == role {
						allowed = true
						break
					}
				}
				if !allowed {
					http.Error(w, "Forbidden", http.StatusForbidden)
					return
				}
			}

			// ✅ Inject user into context
			ctx := context.WithValue(r.Context(), m.cnf.Auth_ctx_key, usr)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
