package user

import (
	"net/http"

	"github.com/gowalillah/ecommerce/rest/middleware"
	"github.com/gowalillah/ecommerce/userRoles"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	// resgister
	mux.Handle(
		"POST /auth/register",
		manager.With(
			http.HandlerFunc(
				h.CreateUser),
		),
	)

	// login
	mux.Handle(
		"POST /auth/login",
		manager.With(
			http.HandlerFunc(
				h.Login,
			),
		),
	)

	// role change -> only can super-admin
	mux.Handle(
		"PUT /auth/status-change",
		manager.With(
			http.HandlerFunc(
				h.ChangeUserRole,
			),
			h.middleware.AuthenticateJwt(userRoles.SuperAdmin),
		),
	)

	// get all users -> only access admin and super-admin
	mux.Handle(
		"GET /users",
		manager.With(
			http.HandlerFunc(
				h.GetUsers,
			),
			h.middleware.AuthenticateJwt(userRoles.Admin, userRoles.SuperAdmin),
		),
	)

	// single user route
	mux.Handle(
		"GET /users/{id}",
		manager.With(
			http.HandlerFunc(
				h.GetUser,
			),
		),
	)

}
