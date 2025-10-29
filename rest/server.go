package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gowalillah/ecommerce/config"
	"github.com/gowalillah/ecommerce/rest/handlers/cart"
	"github.com/gowalillah/ecommerce/rest/handlers/category"
	"github.com/gowalillah/ecommerce/rest/handlers/product"
	"github.com/gowalillah/ecommerce/rest/handlers/user"
	"github.com/gowalillah/ecommerce/rest/middleware"
)

type Server struct {
	cnf             *config.Config
	userHandler     *user.Handler
	categoryHandler *category.Handler
	productHandler  *product.Handler
	cartHandler     *cart.CartHandler
}

func NewServer(
	cnf *config.Config,
	userHandler *user.Handler,
	categoryHandler *category.Handler,
	productHandler *product.Handler,
	cartHandler *cart.CartHandler,

) *Server {
	return &Server{
		cnf:             cnf,
		userHandler:     userHandler,
		categoryHandler: categoryHandler,
		productHandler:  productHandler,
		cartHandler:     cartHandler,
	}
}

func (server *Server) Start() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Logger,
	)

	// Wrap the mux with the middleware
	wrappedMux := manager.WrapMux(mux)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.categoryHandler.RegisterRoutes(mux, manager)
	server.cartHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort) // type casting (int64 to string)

	fmt.Println("Server running PORT:", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error from server", err)
		os.Exit(1)
	}
}
