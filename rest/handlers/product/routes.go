package product

import (
	"net/http"

	"github.com/gowalillah/ecommerce/rest/middleware"
	"github.com/gowalillah/ecommerce/userRoles"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(
				h.GetProducts),
		),
	)

	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(
				h.GetProduct),
		),
	)

	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(
				h.CreateProduct,
			),
			h.middlewares.AuthenticateJwt(userRoles.Admin, userRoles.SuperAdmin),
		),
	)

	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(
				h.UpdateProduct,
			),
			h.middlewares.AuthenticateJwt(userRoles.Admin, userRoles.SuperAdmin),
		),
	)

	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(
				h.DeleteProduct,
			),
			h.middlewares.AuthenticateJwt(userRoles.Admin, userRoles.SuperAdmin),
		),
	)

}
