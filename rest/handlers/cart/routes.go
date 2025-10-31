package cart

import (
	"net/http"

	"github.com/gowalillah/ecommerce/rest/middleware"
)

func (h *CartHandler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"POST /carts",
		manager.With(
			http.HandlerFunc(
				h.CreateCart),
		),
	)

	mux.Handle(
		"GET /carts/{id}",
		manager.With(
			http.HandlerFunc(
				h.GetCart),
		),
	)

	mux.Handle(
		"GET /carts",
		manager.With(
			http.HandlerFunc(
				h.GetCarts,
			),
		),
	)

	mux.Handle(
		"PUT /carts/{id}",
		manager.With(
			http.HandlerFunc(
				h.UpdateCart,
			),
		),
	)

	mux.Handle(
		"DELETE /carts/{id}",
		manager.With(
			http.HandlerFunc(
				h.DeleteCart,
			),
		),
	)

}
