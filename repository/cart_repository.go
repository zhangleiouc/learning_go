package repository

import (
	"context"
	"database/sql"
	"fmt"

	"learning_go/domain"
)

type cartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) domain.CartRepository {
	return &cartRepository{
		db: db,
	}
}

func (cr *cartRepository) AddOrUpdate(ctx context.Context, userID, productID, quantity int64) (int64, error) {
	// First, check if the cart item already exists
	checkQuery := fmt.Sprintf(`
		SELECT id, quantity
		FROM %s
		WHERE user_id = ? AND product_id = ?
		LIMIT 1
	`, domain.TableCart)

	var existingID int64
	var existingQuantity int64
	err := cr.db.QueryRowContext(ctx, checkQuery, userID, productID).Scan(&existingID, &existingQuantity)

	if err != nil && err != sql.ErrNoRows {
		return 0, fmt.Errorf("failed to check existing cart item: %w", err)
	}

	if err == sql.ErrNoRows {
		// Insert new cart item
		insertQuery := fmt.Sprintf(`
			INSERT INTO %s (user_id, product_id, quantity)
			VALUES (?, ?, ?)
		`, domain.TableCart)

		result, err := cr.db.ExecContext(ctx, insertQuery, userID, productID, quantity)
		if err != nil {
			return 0, fmt.Errorf("failed to add cart item: %w", err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("failed to get last insert id: %w", err)
		}

		return id, nil
	}

	// Update existing cart item quantity
	updateQuery := fmt.Sprintf(`
		UPDATE %s
		SET quantity = quantity + ?
		WHERE id = ?
	`, domain.TableCart)

	_, err = cr.db.ExecContext(ctx, updateQuery, quantity, existingID)
	if err != nil {
		return 0, fmt.Errorf("failed to update cart item: %w", err)
	}

	return existingID, nil
}

func (cr *cartRepository) GetByUserID(ctx context.Context, userID int64) ([]*domain.CartItem, error) {
	query := fmt.Sprintf(`
		SELECT 
			c.id, c.user_id, c.product_id, c.quantity, c.created_at,
			p.id, p.product_name
		FROM %s c
		LEFT JOIN %s p ON c.product_id = p.id
		WHERE c.user_id = ?
		ORDER BY c.id DESC
	`, domain.TableCart, domain.TableProduct)

	rows, err := cr.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query cart items: %w", err)
	}
	defer rows.Close()

	var cartItems []*domain.CartItem
	for rows.Next() {
		cartItem := &domain.CartItem{
			Cart:    &domain.Cart{},
			Product: &domain.Product{},
		}
		var createdAt sql.NullString
		var productName sql.NullString

		err := rows.Scan(
			&cartItem.Cart.ID,
			&cartItem.Cart.UserID,
			&cartItem.Cart.ProductID,
			&cartItem.Cart.Quantity,
			&createdAt,
			&cartItem.Product.ID,
			&productName,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cart item: %w", err)
		}

		if createdAt.Valid {
			cartItem.Cart.CreatedAt = &createdAt.String
		}

		if productName.Valid {
			cartItem.Product.Name = &productName.String
		}

		// If product ID is 0 (NULL from LEFT JOIN), set Product to nil
		if cartItem.Product.ID == 0 {
			cartItem.Product = nil
		}

		cartItems = append(cartItems, cartItem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating cart items: %w", err)
	}

	return cartItems, nil
}

func (cr *cartRepository) Delete(ctx context.Context, cartID int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = ?
	`, domain.TableCart)

	result, err := cr.db.ExecContext(ctx, query, cartID)
	if err != nil {
		return fmt.Errorf("failed to delete cart item: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("cart item not found")
	}

	return nil
}
