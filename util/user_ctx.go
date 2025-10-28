package util

import (
	"net/http"

	"github.com/gowalillah/ecommerce/config"
)

func GetUserFromContext(r *http.Request, cnf *config.Config) (*Payload, bool) {
	user, ok := r.Context().Value(cnf.Auth_ctx_key).(*Payload)
	return user, ok
}
