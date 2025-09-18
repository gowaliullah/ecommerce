package main

import (
	"fmt"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my ecommerce......!")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", welcome)
}
