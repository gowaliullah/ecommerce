package middleware

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gowalillah/ecommerce/util"
)

func (m *Middlewares) AuthenticateJwt(next http.Handler) http.Handler {
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

		jswt_token := headerArr[1]

		isVerify, err := util.Verify(jswt_token, m.cnf.JwtSecretKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !isVerify {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		usr, err := util.DecodeToken(jswt_token, m.cnf.JwtSecretKey)
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "failed to verify")
		}

		ctx := context.WithValue(r.Context(), m.cnf.Auth_ctx_key, usr)

		next.ServeHTTP(w, r.WithContext(ctx))

		// accessToken := headerArr[1]
		// tokenParts := strings.Split(accessToken, ".")

		// if len(tokenParts) != 3 {
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// 	return
		// }

		// jwtHeader := tokenParts[0]
		// jwtPayload := tokenParts[1]
		// signature := tokenParts[2]

		// message := jwtHeader + "." + jwtPayload

		// byteArrSecret := []byte(m.cnf.JwtSecretKey)
		// byteArrMsg := []byte(message)

		// h := hmac.New(sha256.New, byteArrSecret)
		// h.Write(byteArrMsg)

		// hash := h.Sum(nil)
		// newSignature := base64UrlEncode(hash)

		// if newSignature != signature {
		// 	http.Error(w, "Unauthorized..! tui hacker", http.StatusUnauthorized)
		// 	return
		// }

		// next.ServeHTTP(w, r)
	})
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
