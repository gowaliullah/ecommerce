package user

import (
	"github.com/gowalillah/ecommerce/domain"
	usrHandler "github.com/gowalillah/ecommerce/rest/handlers/user"
)

type Service interface {
	usrHandler.Service
}

type UserRepo interface {
	Create(u domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
	Get(id string) (*domain.User, error)
	List(limit, page int) ([]*domain.User, error)
	Count() (int64, error)
	Update(u domain.User) (*domain.User, error)
	UpdateRole(id int, newRole string) error
	Delete(id string) error
}
