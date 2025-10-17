package cart

import "github.com/gowalillah/ecommerce/domain"

type service struct {
	cartRepo CartRepo
}

func NewService(cartRepo CartRepo) Service {
	return &service{
		cartRepo: cartRepo,
	}
}

func (svc *service) Create(c domain.Cart) (*domain.Cart, error) {
	return svc.cartRepo.Create(c)
}

func (svc *service) Get(id int) (*domain.Cart, error) {
	return svc.cartRepo.Get(id)
}

func (svc *service) List() ([]*domain.Cart, error) {
	return svc.cartRepo.List()
}

func (svc *service) Update(c domain.Cart) (*domain.Cart, error) {
	return svc.cartRepo.Update(c)
}

func (svc *service) Delete(id int) error {
	return svc.cartRepo.Delete(id)
}
