package repository

import (
	"context"
	"database/sql"
	"fmt"

	"learning_go/domain"
)

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) domain.OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (or *orderRepository) GetByID(c context.Context, id int64) (*domain.Order, error) {
	query := fmt.Sprintf(`
		SELECT id, order_code, customer_id
		FROM %s
		WHERE id = ?
	`, domain.TableOrder)

	order := &domain.Order{}
	var orderNo sql.NullString
	var customerID sql.NullInt64

	err := or.db.QueryRowContext(c, query, id).Scan(
		&order.ID,
		&orderNo,
		&customerID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("order not found")
		}
		return nil, err
	}

	if orderNo.Valid {
		order.OrderCode = &orderNo.String
	}
	if customerID.Valid {
		order.CustomerID = &customerID.Int64
	}
	//if totalAmount.Valid {
	//	order.TotalAmount = &totalAmount.String
	//}
	//if status.Valid {
	//	order.Status = &status.String
	//}
	//if createdAt.Valid {
	//	order.CreatedAt = &createdAt.String
	//}
	//if updatedAt.Valid {
	//	order.UpdatedAt = &updatedAt.String
	//}

	return order, nil
}

func (or *orderRepository) Create(c context.Context, order *domain.Order) (int64, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (order_no, customer_id)
		VALUES (?, ?)
	`, domain.TableOrder)

	result, err := or.db.ExecContext(c, query, order.OrderCode, order.CustomerID)
	if err != nil {
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, nil
}
