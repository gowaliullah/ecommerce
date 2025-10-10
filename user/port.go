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
	Get(id int) (*domain.User, error)
	List() ([]*domain.User, error)
	Count() (int64, error)
	Update(u domain.User) (*domain.User, error)
	Delete(id int) error
}
