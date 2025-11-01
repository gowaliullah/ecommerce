package order

import (
	"net/http"

	"github.com/gowalillah/ecommerce/rest/middleware"
)

func (h *OrderHandler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"GET /orders",
		manager.With(
			http.HandlerFunc(
				h.CreateOrder),
		),
	)

	mux.Handle(
		"GET /orders/{id}",
		manager.With(
			http.HandlerFunc(
				h.GetOrder),
		),
	)

	mux.Handle(
		"POST /orders",
		manager.With(
			http.HandlerFunc(
				h.GetOrders,
			),
		),
	)

	mux.Handle(
		"PUT /orders/{id}",
		manager.With(
			http.HandlerFunc(
				h.UpdateOrder,
			),
		),
	)

	mux.Handle(
		"DELETE /orders/{id}",
		manager.With(
			http.HandlerFunc(
				h.DeleteOrder,
			),
		),
	)

}
