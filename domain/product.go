package domain

import "context"

const (
	TableProduct = "oms_product"
)

// Product represents the minimal fields we expose for product search.
// Extend this struct if the oms_product table contains more columns that need to be returned.
type Product struct {
	ID   int64   `json:"id" db:"id"`
	Name *string `json:"name,omitempty" db:"product_name"`
}

type ProductRepository interface {
	// SearchByName performs a fuzzy search on product name.
	SearchByName(ctx context.Context, keyword string) ([]*Product, error)
}

type ProductUsecase interface {
	SearchByName(ctx context.Context, keyword string) ([]*Product, error)
}
