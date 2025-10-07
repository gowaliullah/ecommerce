package product

import (
	"github.com/gowalillah/ecommerce/domain"
	prdHandler "github.com/gowalillah/ecommerce/rest/handlers/product"
)

type Service interface {
	prdHandler.Service
}

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Count() (int64, error)
	Update(p domain.Product) (*domain.Product, error)
	Delete(id int) error
}
