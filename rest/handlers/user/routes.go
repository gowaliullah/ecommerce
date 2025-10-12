package user

import (
	"net/http"

	"github.com/gowalillah/ecommerce/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"POST /users",
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
