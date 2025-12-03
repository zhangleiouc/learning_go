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
		SELECT id, order_code, customer_id, status
		FROM %s
		WHERE id = ?
	`, domain.TableOrder)

	order := &domain.Order{}
	var orderNo sql.NullString
	var customerID sql.NullInt64
	var status sql.NullString

	err := or.db.QueryRowContext(c, query, id).Scan(
		&order.ID,
		&orderNo,
		&customerID,
		&status,
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
	if status.Valid {
		order.Status = &status.String
	}
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
		INSERT INTO %s (order_no, customer_id, status)
		VALUES (?, ?, ?)
	`, domain.TableOrder)

	result, err := or.db.ExecContext(c, query, order.OrderCode, order.CustomerID, order.Status)
	if err != nil {
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, nil
}

func (or *orderRepository) GetByCustomerID(c context.Context, customerID int64) ([]*domain.Order, error) {
	query := fmt.Sprintf(`
		SELECT id, order_code, customer_id, status
		FROM %s
		WHERE customer_id = ?
		ORDER BY id DESC
	`, domain.TableOrder)

	rows, err := or.db.QueryContext(c, query, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to query orders: %w", err)
	}
	defer rows.Close()

	var orders []*domain.Order
	for rows.Next() {
		order := &domain.Order{}
		var orderNo sql.NullString
		var customerID sql.NullInt64
		var status sql.NullString

		err := rows.Scan(
			&order.ID,
			&orderNo,
			&customerID,
			&status,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}

		if orderNo.Valid {
			order.OrderCode = &orderNo.String
		}
		if customerID.Valid {
			order.CustomerID = &customerID.Int64
		}
		if status.Valid {
			order.Status = &status.String
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating orders: %w", err)
	}

	return orders, nil
}

// UpdateStatus 更新订单状态
func (or *orderRepository) UpdateStatus(c context.Context, id int64, status string) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET status = ?
		WHERE id = ?
	`, domain.TableOrder)

	result, err := or.db.ExecContext(c, query, status, id)
	if err != nil {
		return fmt.Errorf("failed to update order status: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected when updating order status: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("order not found")
	}

	return nil
}
