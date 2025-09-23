package cmd

import (
	"fmt"
	"net/http"

	"github.com/gowalillah/ecommerce/middlewares"
	"github.com/gowalillah/ecommerce/rest/handlers/products"
)

func Serve() {
	mux := http.NewServeMux()

	// mux.Handle("/products", http.HandlerFunc(getProducts))

	mux.Handle("GET products", http.HandlerFunc(products.GetProducts))

	mux.Handle("/create-product", http.HandlerFunc(products.CreateProduct))

	fmt.Println("Server running port on:8080")
	err := http.ListenAndServe(":8080", middlewares.GlobalRouter(mux))
	if err != nil {
		fmt.Println("Error from server", err)
	}
}
