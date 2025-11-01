package order

import (
	"github.com/gowalillah/ecommerce/config"
	"github.com/gowalillah/ecommerce/rest/middleware"
)

type OrderHandler struct {
	middlewares middleware.Middlewares
	cnf         *config.Config
	svc         Service
}

func NewOrderHandler(middlewares *middleware.Middlewares, cnf *config.Config, svc Service) *OrderHandler {
	return &OrderHandler{
		middlewares: *middlewares,
		cnf:         cnf,
		svc:         svc,
	}
}
