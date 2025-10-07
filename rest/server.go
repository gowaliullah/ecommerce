package rest

import (
	"fmt"
	"net/http"

	"github.com/gowalillah/ecommerce/config"
	"github.com/gowalillah/ecommerce/middleware"
)

type Server struct {
	cnf *config.Config
}

func (server *Server) Start() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.Logger, middleware.Logger)

	// initRoutes(mux, manager)

	fmt.Println("Server running port on:8080")
	err := http.ListenAndServe(":8080", middleware.GlobalRouter(mux))
	if err != nil {
		fmt.Println("Error from server", err)
	}
}
