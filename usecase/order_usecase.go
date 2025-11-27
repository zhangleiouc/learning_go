package usecase

import (
	"context"
	"time"

	"learning_go/domain"
)

type orderUsecase struct {
	orderRepository domain.OrderRepository
	contextTimeout  time.Duration
}

func NewOrderUsecase(orderRepository domain.OrderRepository, timeout time.Duration) domain.OrderUsecase {
	return &orderUsecase{
		orderRepository: orderRepository,
		contextTimeout:  timeout,
	}
}

func (ou *orderUsecase) GetByID(c context.Context, id int64) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.GetByID(ctx, id)
}

func (ou *orderUsecase) Create(c context.Context, request *domain.CreateOrderRequest) (int64, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()

	order := &domain.Order{
		OrderCode:  &request.OrderCode,
		CustomerID: &request.CustomerID,
	}

	return ou.orderRepository.Create(ctx, order)
}
