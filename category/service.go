package category

import "github.com/gowalillah/ecommerce/domain"

type service struct {
	catgRepo CategoryRepo
}

func NewService(catgRepo CategoryRepo) Service {
	return &service{
		catgRepo: catgRepo,
	}
}

func (svc *service) Create(c domain.Category) (*domain.Category, error) {
	return svc.catgRepo.Create(c)
}

func (svc *service) Get(id int) (*domain.Category, error) {
	return svc.catgRepo.Get(id)
}

func (svc *service) List() ([]*domain.Category, error) {
	return svc.catgRepo.List()
}

func (svc *service) Update(c domain.Category) (*domain.Category, error) {
	return svc.catgRepo.Update(c)
}

func (svc *service) Delete(id int) error {
	return svc.catgRepo.Delete(id)
}
