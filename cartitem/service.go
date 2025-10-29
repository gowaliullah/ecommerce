package cartitem

import "github.com/gowalillah/ecommerce/domain"

type service struct {
	cartItemRepo CartItemRepo
}

func NewService(cartItemRepo CartItemRepo) Service {
	return &service{
		cartItemRepo: cartItemRepo,
	}
}

func (svc *service) Create(c domain.CartItem) (*domain.CartItem, error) {
	return svc.cartItemRepo.Create(c)
}

func (svc *service) Get(id int) (*domain.CartItem, error) {
	return svc.cartItemRepo.Get(id)
}

func (svc *service) List() ([]*domain.CartItem, error) {
	return svc.cartItemRepo.List()
}

func (svc *service) Update(c domain.CartItem) (*domain.CartItem, error) {
	return svc.cartItemRepo.Update(c)
}

func (svc *service) Delete(id int) error {
	return svc.cartItemRepo.Delete(id)
}
