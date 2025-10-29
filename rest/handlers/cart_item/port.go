package cartitem

import "github.com/gowalillah/ecommerce/domain"

type Service interface {
	Create(cart domain.Cart) (*domain.Cart, error)
	Get(id int) (*domain.Cart, error)
	List() ([]*domain.Cart, error)
	Update(cart domain.Cart) (*domain.Cart, error)
	Delete(id int) error
}
