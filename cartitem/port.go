package cartitem

import (
	"github.com/gowalillah/ecommerce/domain"
	cartItemHandler "github.com/gowalillah/ecommerce/rest/handlers/cartitem"
)

type Service interface {
	cartItemHandler.Service
}

type CartItemRepo interface {
	Create(cart domain.CartItem) (*domain.CartItem, error)
	Get(id int) (*domain.CartItem, error)
	List() ([]*domain.CartItem, error)
	Update(cartItem domain.CartItem) (*domain.CartItem, error)
	Delete(id int) error
}
