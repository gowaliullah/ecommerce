package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gowalillah/ecommerce/util"
)

func (m *Middlewares) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		AuthenticationHeader := r.Header.Get("Authorization")
		if AuthenticationHeader == "" {
			http.Error(w, "Unauthorized:", http.StatusUnauthorized)
			return
		}

		//split header and grep the token
		headerArr := strings.Split(AuthenticationHeader, " ")
		if len(headerArr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwt_token := headerArr[1]

		// verify token
		isVerified, err := util.Verify(jwt_token, m.cnf.JwtSecretKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !isVerified {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user, err := util.DecodeToken(jwt_token, m.cnf.JwtSecretKey)
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "failed to verify")
		}

		ctx := context.WithValue(r.Context(), m.cnf.Auth_ctx_key, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middlewares) AuthenticateJwt33(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// 1️⃣ Check Authorization header
			header := r.Header.Get("Authorization")
			if header == "" {
				http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(header, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				http.Error(w, "Unauthorized: invalid format", http.StatusUnauthorized)
				return
			}

			token := parts[1]

			// 2️⃣ Verify token validity
			ok, err := util.Verify(token, m.cnf.JwtSecretKey)
			if err != nil || !ok {
				http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
				return
			}

			// 3️⃣ Decode token payload
			usr, err := util.DecodeToken(token, m.cnf.JwtSecretKey)
			if err != nil {
				http.Error(w, "Unauthorized: failed to decode token", http.StatusUnauthorized)
				return
			}

			// 4️⃣ Role-based check (using if conditions)
			if len(allowedRoles) > 0 {
				// Example: route protected by SuperAdmin only
				if len(allowedRoles) == 1 && allowedRoles[0] == "super-admin" {
					if usr.Role != "super-admin" {
						http.Error(w, "Forbidden: super-admin access required", http.StatusForbidden)
						return
					}
				}

				// Example: route protected by Admin only
				if len(allowedRoles) == 1 && allowedRoles[0] == "admin" {
					if usr.Role != "admin" {
						http.Error(w, "Forbidden: admin access required", http.StatusForbidden)
						return
					}
				}

				// Example: route protected by Admin or SuperAdmin
				if len(allowedRoles) > 1 {
					allowed := false
					for _, role := range allowedRoles {
						if usr.Role == role {
							allowed = true
							break
						}
					}
					if !allowed {
						http.Error(w, "Forbidden: insufficient role", http.StatusForbidden)
						return
					}
				}
			}

			// 5️⃣ Add user to context
			ctx := context.WithValue(r.Context(), m.cnf.Auth_ctx_key, usr)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
