package user

import (
	"github.com/gowalillah/ecommerce/config"
	"github.com/gowalillah/ecommerce/rest/middleware"
)

type Handler struct {
	middleware middleware.Middlewares
	cnf        *config.Config
	svc        Service
}

func NewHandler(middlewares *middleware.Middlewares, cnf *config.Config, svc Service) *Handler {
	return &Handler{
		middleware: *middlewares,
		cnf:        cnf,
		svc:        svc,
	}
}
