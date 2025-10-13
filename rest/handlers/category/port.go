package category

import "github.com/gowalillah/ecommerce/domain"

type Service interface {
	Create(data domain.Category) (*domain.Category, error)
	Get(id int) (*domain.Category, error)
	List() ([]*domain.Category, error)
	Update(data domain.Category) (*domain.Category, error)
	Delete(id int) error
}
