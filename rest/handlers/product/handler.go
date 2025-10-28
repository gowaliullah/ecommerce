package product

import (
	"github.com/gowalillah/ecommerce/config"
	"github.com/gowalillah/ecommerce/rest/middleware"
)

type Handler struct {
	middlewares middleware.Middlewares
	cnf         *config.Config
	svc         Service
}

func NewHandler(middlewares *middleware.Middlewares, cnf *config.Config, svc Service) *Handler {
	return &Handler{
		middlewares: *middlewares,
		cnf:         cnf,
		svc:         svc,
	}
}
