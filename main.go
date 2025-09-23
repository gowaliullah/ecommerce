package main

import (
	"fmt"
	"net/http"

	"github.com/gowalillah/ecommerce/cmd"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my ecommerce......!")
}

func main() {
	cmd.Serve()
}
