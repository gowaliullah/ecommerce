package cart

import (
	"github.com/gowalillah/ecommerce/domain"
	cartHandler "github.com/gowalillah/ecommerce/rest/handlers/cart"
)

type Service interface {
	cartHandler.Service
}

type CartRepo interface {
	Create(cart domain.Cart) (*domain.Cart, error)
	Get(id int) (*domain.Cart, error)
	List() ([]*domain.Cart, error)
	Update(cart domain.Cart) (*domain.Cart, error)
	Delete(id int) error
}
