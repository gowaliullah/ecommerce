package order

import (
	"github.com/gowalillah/ecommerce/domain"
	OrderHandler "github.com/gowalillah/ecommerce/rest/handlers/order"
)

type Service interface {
	OrderHandler.Service
}

type OrderRepo interface {
	Create(cart domain.Order) (*domain.Order, error)
	Get(id int) (*domain.Order, error)
	List() ([]*domain.Order, error)
	Update(order domain.Order) (*domain.Order, error)
	Delete(id int) error
}
