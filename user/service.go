package user

import (
	"errors"

	"github.com/gowalillah/ecommerce/domain"
)

type service struct {
	usrRepo UserRepo
}

func NewService(usrRepo UserRepo) Service {
	return &service{
		usrRepo: usrRepo,
	}
}

func (svc *service) Create(u domain.User) (*domain.User, error) {
	user, err := svc.usrRepo.Create(u)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return user, nil
}

func (svc *service) Find(email, pass string) (*domain.User, error) {
	user, err := svc.usrRepo.Find(email, pass)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (svc *service) Count() (int64, error) {
	return svc.usrRepo.Count()
}

func (svc *service) List(limit, page int) ([]*domain.User, error) {
	users, err := svc.usrRepo.List(limit, page)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (svc *service) Get(id string) (*domain.User, error) {
	return svc.usrRepo.Get(id)
}

func (svc *service) Update(u domain.User) (*domain.User, error) {
	return svc.usrRepo.Update(u)
}

func (svc *service) Delete(id string) error {
	return svc.usrRepo.Delete(id)
}
