package cartitem

import (
	"github.com/gowalillah/ecommerce/config"
	"github.com/gowalillah/ecommerce/rest/middleware"
)

type CartItemHandler struct {
	middlewares middleware.Middlewares
	cnf         *config.Config
	svc         Service
}

func NewCartItemHandler(middlewares *middleware.Middlewares, cnf *config.Config, svc Service) *CartItemHandler {
	return &CartItemHandler{
		middlewares: *middlewares,
		cnf:         cnf,
		svc:         svc,
	}
}
