package category

import (
	"net/http"

	"github.com/gowalillah/ecommerce/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"GET /categories",
		manager.With(
			http.HandlerFunc(
				h.GetCategories),
		),
	)

	mux.Handle(
		"GET /categories/{id}",
		manager.With(
			http.HandlerFunc(
				h.GetCategory),
		),
	)

	mux.Handle(
		"POST /categories",
		manager.With(
			http.HandlerFunc(
				h.CreateCategory,
			),
		),
	)

	mux.Handle(
		"PUT /categories/{id}",
		manager.With(
			http.HandlerFunc(
				h.UpdateCategory,
			),
		),
	)

	mux.Handle(
		"DELETE /categories/{id}",
		manager.With(
			http.HandlerFunc(
				h.DeleteCategory,
			),
		),
	)

}
