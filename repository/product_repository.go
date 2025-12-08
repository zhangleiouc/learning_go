package repository

import (
	"context"
	"database/sql"
	"fmt"

	"learning_go/domain"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (pr *productRepository) SearchByName(ctx context.Context, keyword string) ([]*domain.Product, error) {
	query := fmt.Sprintf(`
		SELECT id, product_name
		FROM %s
		WHERE product_name LIKE ?
		ORDER BY id DESC
	`, domain.TableProduct)

	rows, err := pr.db.QueryContext(ctx, query, "%"+keyword+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to search products: %w", err)
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		product := &domain.Product{}
		var name sql.NullString

		if err := rows.Scan(&product.ID, &name); err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}

		if name.Valid {
			product.Name = &name.String
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating products: %w", err)
	}

	return products, nil
}
