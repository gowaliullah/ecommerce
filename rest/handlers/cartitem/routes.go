package cartitem

import (
	"net/http"

	"github.com/gowalillah/ecommerce/rest/middleware"
)

func (h *CartItemHandler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"GET /cart-items",
		manager.With(
			http.HandlerFunc(
				h.CreateCartItem),
		),
	)

	mux.Handle(
		"GET /cart-items/{id}",
		manager.With(
			http.HandlerFunc(
				h.GetCartItem),
		),
	)

	mux.Handle(
		"POST /cart-items",
		manager.With(
			http.HandlerFunc(
				h.GetCartItems,
			),
		),
	)

	mux.Handle(
		"PUT /cart-items/{id}",
		manager.With(
			http.HandlerFunc(
				h.UpdateCartItem,
			),
		),
	)

	mux.Handle(
		"DELETE /cart-items/{id}",
		manager.With(
			http.HandlerFunc(
				h.DeleteCartItem,
			),
		),
	)

}
