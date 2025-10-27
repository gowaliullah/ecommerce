package user

import "github.com/gowalillah/ecommerce/domain"

type Service interface {
	Create(user domain.User) (*domain.User, error)
	Find(Email, Password string) (*domain.User, error)
	List(limit, page int) ([]*domain.User, error)
	Get(id string) (*domain.User, error)
	Update(user domain.User) (*domain.User, error)
	Delete(id string) error
}
