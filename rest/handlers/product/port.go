package product

import "github.com/gowalillah/ecommerce/domain"

type Port interface {
	Create(domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Count() (int64, error)
	Update(domain.Product) (*domain.Product, error)
	Delete(id int) error
}
