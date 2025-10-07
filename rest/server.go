package rest

import (
	"fmt"
	"net/http"

	"github.com/gowalillah/ecommerce/config"
	"github.com/gowalillah/ecommerce/rest/middleware"
)

type Server struct {
	cnf *config.Config
}

func NewServer(
	cnf *config.Config,
) *Server {
	return &Server{
		cnf: cnf,
	}
}

func (server *Server) Start() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(middleware.Logger)

	// Wrap the mux with the middleware
	wrappedMux := manager.WrapMux(mux)

	// initRoutes(mux, manager)

	fmt.Println("Server running port on:8080")
	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error from server", err)
	}
}
