package domain

import (
	"context"
)

const (
	TableOrder = "oms_order"
)

type Order struct {
	ID         int64   `json:"id" db:"id"`
	OrderCode  *string `json:"order_code,omitempty" db:"order_no"`
	CustomerID *int64  `json:"customer_id,omitempty" db:"customer_id"`
	//TotalAmount *string `json:"total_amount,omitempty" db:"total_amount"`
	//Status *string `json:"status,omitempty" db:"status"`
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
}

type OrderUsecase interface {
	GetByID(c context.Context, id int64) (*Order, error)
	Create(c context.Context, request *CreateOrderRequest) (int64, error)
	GetByCustomerID(c context.Context, customerID int64) ([]*Order, error)
}
