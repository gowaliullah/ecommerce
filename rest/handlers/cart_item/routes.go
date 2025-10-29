package cartitem

import (
	"net/http"

	"github.com/gowalillah/ecommerce/rest/middleware"
)

func (h *CartItemHandler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"GET /carts",
		manager.With(
			http.HandlerFunc(
				h.CreateCartItem),
		),
	)

	mux.Handle(
		"GET /carts/{id}",
		manager.With(
			http.HandlerFunc(
				h.GetCartItem),
		),
	)

	mux.Handle(
		"POST /carts",
		manager.With(
			http.HandlerFunc(
				h.GetCartItems,
			),
		),
	)

	mux.Handle(
		"PUT /carts/{id}",
		manager.With(
			http.HandlerFunc(
				h.UpdateCartItem,
			),
		),
	)

	mux.Handle(
		"DELETE /carts/{id}",
		manager.With(
			http.HandlerFunc(
				h.DeleteCartItem,
			),
		),
	)

}
