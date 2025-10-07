package product

import (
	"github.com/gowalillah/ecommerce/domain"
)

type service struct {
	prdRepo ProductRepo
}

func NewService(prdRepo ProductRepo) Service {
	return &service{
		prdRepo: prdRepo,
	}
}

func (svc *service) Create(prd domain.Product) (*domain.Product, error) {
	return svc.prdRepo.Create(prd)
}

func (svc *service) Count() (int64, error) {
	return svc.prdRepo.Count()
}

func (svc *service) List() ([]*domain.Product, error) {
	return svc.prdRepo.List()
}

func (svc *service) Get(id int) (*domain.Product, error) {
	return svc.prdRepo.Get(id)
}

func (svc *service) Update(prd domain.Product) (*domain.Product, error) {
	return svc.prdRepo.Update(prd)
}

func (svc *service) Delete(id int) error {
	return svc.prdRepo.Delete(id)
}
