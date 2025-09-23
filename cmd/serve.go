package cmd

import (
	"fmt"
	"net/http"

	"github.com/gowalillah/ecommerce/middlewares"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(welcome))

	// mux.Handle("/products", http.HandlerFunc(getProducts))

	mux.Handle("GET products", http.HandlerFunc(getProducts))

	mux.Handle("/create-product", http.HandlerFunc(createProduct))

	fmt.Println("Server running port on:8080")
	err := http.ListenAndServe(":8080", middlewares.GlobalRouter(mux))
	if err != nil {
		fmt.Println("Error from server", err)
	}
}
