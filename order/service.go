package order

import "github.com/gowalillah/ecommerce/domain"

type service struct {
	orderRepo OrderRepo
}

func NewService(orderRepo OrderRepo) Service {
	return &service{
		orderRepo: orderRepo,
	}
}

func (svc *service) Create(c domain.Order) (*domain.Order, error) {
	return svc.orderRepo.Create(c)
}

func (svc *service) Get(id int) (*domain.Order, error) {
	return svc.orderRepo.Get(id)
}

func (svc *service) List() ([]*domain.Order, error) {
	return svc.orderRepo.List()
}

func (svc *service) Update(c domain.Order) (*domain.Order, error) {
	return svc.orderRepo.Update(c)
}

func (svc *service) Delete(id int) error {
	return svc.orderRepo.Delete(id)
}
