package user

import (
	"net/http"

	"github.com/gowalillah/ecommerce/rest/middleware"
	"github.com/gowalillah/ecommerce/userRoles"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"POST /users/register",
		manager.With(
			http.HandlerFunc(
				h.CreateUser),
		),
	)

	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(
				h.Login,
			),
		),
	)

	mux.Handle(
		"GET /users",
		manager.With(
			http.HandlerFunc(
				h.GetUsers,
			),
			h.middleware.AuthenticateJwt(userRoles.Admin),
		),
	)

	mux.Handle(
		"GET /users/{id}",
		manager.With(
			http.HandlerFunc(
				h.GetUser,
			),
		),
	)

}
