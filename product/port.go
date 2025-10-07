package product

import "github.com/gowalillah/ecommerce/domain"

type Service interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Count() (int64, error)
	Update(p domain.Product) (*domain.Product, error)
	Delete(id int) error
}
