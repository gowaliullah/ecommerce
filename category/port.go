package category

import (
	"github.com/gowalillah/ecommerce/domain"
	catHandler "github.com/gowalillah/ecommerce/rest/handlers/category"
)

type Service interface {
	catHandler.Service
}

type CategoryRepo interface {
	Create(data domain.Category) (*domain.Category, error)
	Get(id int) (*domain.Category, error)
	List() ([]*domain.Category, error)
	Update(data domain.Category) (*domain.Category, error)
	Delete(id int) error
}
