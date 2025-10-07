package user

import "github.com/gowalillah/ecommerce/domain"

type Service interface {
	Create(user domain.User) (*domain.User, error)
	List() ([]*domain.User, error)
	Get(id int) (*domain.User, error)
	Update(user domain.User) (*domain.User, error)
	Delete(id int) error
}
