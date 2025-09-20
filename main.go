package main

import (
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my ecommerce......!")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", welcome)

	fmt.Println("Server running port on:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error from server", err)
	}
}
