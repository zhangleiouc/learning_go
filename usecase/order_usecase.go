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
		Status:     ptrString(domain.OrderStatusPendingPayment),
	}

	return ou.orderRepository.Create(ctx, order)
}

func (ou *orderUsecase) GetByCustomerID(c context.Context, customerID int64) ([]*domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.GetByCustomerID(ctx, customerID)
}

// MarkAsPaid 将订单从“待付款”变更为“已付款”
// 如果订单状态不是“待付款”，则返回 ErrOrderStatusAlreadyUpdated
func (ou *orderUsecase) MarkAsPaid(c context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()

	order, err := ou.orderRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if order.Status == nil || *order.Status != domain.OrderStatusPendingPayment {
		return domain.ErrOrderStatusAlreadyUpdated
	}

	return ou.orderRepository.UpdateStatus(ctx, id, domain.OrderStatusPaid)
}

func ptrString(s string) *string {
	return &s
}
