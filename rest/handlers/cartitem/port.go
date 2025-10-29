package cartitem

import "github.com/gowalillah/ecommerce/domain"

type Service interface {
	Create(cart domain.CartItem) (*domain.CartItem, error)
	Get(id int) (*domain.CartItem, error)
	List() ([]*domain.CartItem, error)
	Update(cartItem domain.CartItem) (*domain.CartItem, error)
	Delete(id int) error
}
