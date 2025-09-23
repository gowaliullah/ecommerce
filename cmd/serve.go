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

	manager.Use(middleware.Logger, middleware.Logger, middleware.Logger)

	mux := http.NewServeMux()

	mux.Handle("GET /", manager.With(
		http.HandlerFunc(handlers.Welcome),
	))

	mux.Handle("GET /products", http.HandlerFunc(products.GetProducts))

	mux.Handle("/create-product", http.HandlerFunc(products.CreateProduct))

	fmt.Println("Server running port on:8080")
	err := http.ListenAndServe(":8080", middleware.GlobalRouter(mux))
	if err != nil {
		fmt.Println("Error from server", err)
	}
}
