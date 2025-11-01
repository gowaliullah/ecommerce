package order

import "github.com/gowalillah/ecommerce/domain"

type Service interface {
	Create(cart domain.Order) (*domain.Order, error)
	Get(id int) (*domain.Order, error)
	List() ([]*domain.Order, error)
	Update(cartItem domain.Order) (*domain.Order, error)
	Delete(id int) error
}
