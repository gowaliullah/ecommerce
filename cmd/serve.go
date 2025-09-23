package cmd

import (
	"fmt"
	"net/http"

	"github.com/gowalillah/ecommerce/middleware"
	"github.com/gowalillah/ecommerce/rest/handlers"
	"github.com/gowalillah/ecommerce/rest/handlers/products"
)

func Serve() {

	manager := middleware.NewManager()

	mux := http.NewServeMux()

	mux.Handle("GET /", manager.With2(
		http.HandlerFunc(handlers.Welcome),
		middleware.Logger,
		middleware.Logger,
	))

	mux.Handle("GET /products", http.HandlerFunc(products.GetProducts))

	mux.Handle("/create-product", http.HandlerFunc(products.CreateProduct))

	fmt.Println("Server running port on:8080")
	err := http.ListenAndServe(":8080", middleware.GlobalRouter(mux))
	if err != nil {
		fmt.Println("Error from server", err)
	}
}
