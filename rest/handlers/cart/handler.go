package cart

import (
	"github.com/gowalillah/ecommerce/config"
	"github.com/gowalillah/ecommerce/rest/middleware"
)

type CartHandler struct {
	middlewares middleware.Middlewares
	cnf         *config.Config
	svc         Service
}

func NewCartHandler(middlewares *middleware.Middlewares, cnf *config.Config, svc Service) *CartHandler {
	return &CartHandler{
		middlewares: *middlewares,
		cnf:         cnf,
		svc:         svc,
	}
}
