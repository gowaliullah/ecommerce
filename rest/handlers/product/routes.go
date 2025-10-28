package product

import (
	"net/http"

	"github.com/gowalillah/ecommerce/rest/middleware"
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
			h.middlewares.AuthenticateJwt("admin"),
		),
	)

	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(
				h.UpdateProduct,
			),
		),
	)

	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(
				h.DeleteProduct,
			),
		),
	)

}
