package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gowalillah/ecommerce/config"
	"github.com/gowalillah/ecommerce/rest/handlers/product"
	"github.com/gowalillah/ecommerce/rest/middleware"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *product.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
	}
}

func (server *Server) Start() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(middleware.Logger)

	// Wrap the mux with the middleware
	wrappedMux := manager.WrapMux(mux)

	server.productHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort) // type casting (int64 to string)

	fmt.Println("Server running PORT:", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error from server", err)
	}
}
