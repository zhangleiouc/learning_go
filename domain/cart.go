package domain

import "context"

const (
	TableCart = "shopping_cart"
)

// Cart represents a shopping cart item.
type Cart struct {
	ID        int64   `json:"id" db:"id"`
	UserID    int64   `json:"user_id" db:"user_id"`
	ProductID int64   `json:"product_id" db:"product_id"`
	Quantity  int64   `json:"quantity" db:"quantity"`
	CreatedAt *string `json:"created_at,omitempty" db:"created_at"`
}

// CartItem represents a cart item with associated product information.
// Used when browsing cart to include product details.
type CartItem struct {
	Cart    *Cart    `json:"cart"`
	Product *Product `json:"product,omitempty"`
}

// AddToCartRequest represents the request to add a product to cart.
type AddToCartRequest struct {
	UserID    int64 `json:"user_id" binding:"required"`
	ProductID int64 `json:"product_id" binding:"required"`
	Quantity  int64 `json:"quantity" binding:"required,min=1"`
}

type CartRepository interface {
	// AddOrUpdate adds a product to cart or updates quantity if already exists.
	AddOrUpdate(ctx context.Context, userID, productID, quantity int64) (int64, error)
	// GetByUserID returns all cart items for a user with product information.
	GetByUserID(ctx context.Context, userID int64) ([]*CartItem, error)
	// Delete removes a cart item by cart ID.
	Delete(ctx context.Context, cartID int64) error
}

type CartUsecase interface {
	// AddOrUpdate adds a product to cart or updates quantity if already exists.
	AddOrUpdate(ctx context.Context, userID, productID, quantity int64) (int64, error)
	// GetByUserID returns all cart items for a user with product information.
	GetByUserID(ctx context.Context, userID int64) ([]*CartItem, error)
	// Delete removes a cart item by cart ID.
	Delete(ctx context.Context, cartID int64) error
}
