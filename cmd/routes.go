package cmd

import (
	"net/http"

	"github.com/gowalillah/ecommerce/rest/handlers"
	"github.com/gowalillah/ecommerce/rest/handlers/products"
	"github.com/gowalillah/ecommerce/rest/middleware"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /", manager.With(
		http.HandlerFunc(handlers.Welcome),
	))

	mux.Handle("GET /products", http.HandlerFunc(products.GetProducts))

	mux.Handle("/create-product", http.HandlerFunc(products.CreateProduct))

}
