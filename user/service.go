package user

import "github.com/gowalillah/ecommerce/domain"

type service struct {
	usrRepo UserRepo
}

func NewService(usrRepo UserRepo) Service {
	return &service{
		usrRepo: usrRepo,
	}
}

func (svc *service) Create(u domain.User) (*domain.User, error) {
	return svc.usrRepo.Create(u)
}

func (svc *service) Count() (int64, error) {
	return svc.usrRepo.Count()
}

func (svc *service) List() ([]*domain.User, error) {
	return svc.usrRepo.List()
}

func (svc *service) Get(id int) (*domain.User, error) {
	return svc.usrRepo.Get(id)
}

func (svc *service) Update(u domain.User) (*domain.User, error) {
	return svc.usrRepo.Update(u)
}

func (svc *service) Delete(id int) error {
	return svc.usrRepo.Delete(id)
}
