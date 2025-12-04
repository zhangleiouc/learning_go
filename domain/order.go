package domain

import (
	"context"
	"errors"
)

const (
	TableOrder = "oms_order"

	// 订单状态
	OrderStatusPendingPayment = "PENDING_PAYMENT" // 待付款
	OrderStatusPaid           = "PAID"            // 已付款
)

// 业务错误
var (
	// 当订单状态已从待付款变更（例如已付款、已取消等）时，用于告知上层“订单状态已更新，请刷新后尝试”
	ErrOrderStatusAlreadyUpdated = errors.New("order status already updated")
)

type Order struct {
	ID         int64   `json:"id" db:"id"`
	OrderCode  *string `json:"order_code,omitempty" db:"order_no"`
	CustomerID *int64  `json:"customer_id,omitempty" db:"customer_id"`
	Status     *string `json:"status,omitempty" db:"status"`
	//TotalAmount *string `json:"total_amount,omitempty" db:"total_amount"`
	//CreatedAt   *string `json:"created_at,omitempty" db:"created_at"`
	//UpdatedAt   *string `json:"updated_at,omitempty" db:"updated_at"`
	// 可以根据实际表结构添加更多字段
}

type CreateOrderRequest struct {
	OrderCode  string `json:"order_code" binding:"required"`
	CustomerID int64  `json:"customer_id" binding:"required"`
}

type CreateOrderResponse struct {
	OrderID int64 `json:"order_id"`
}

type OrderRepository interface {
	GetByID(c context.Context, id int64) (*Order, error)
	Create(c context.Context, order *Order) (int64, error)
	GetByCustomerID(c context.Context, customerID int64) ([]*Order, error)
	UpdateStatus(c context.Context, id int64, status string) error
}

type OrderUsecase interface {
	GetByID(c context.Context, id int64) (*Order, error)
	Create(c context.Context, request *CreateOrderRequest) (int64, error)
	GetByCustomerID(c context.Context, customerID int64) ([]*Order, error)
	MarkAsPaid(c context.Context, id int64) error
}
