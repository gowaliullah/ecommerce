package product

import "github.com/gowalillah/ecommerce/domain"

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
