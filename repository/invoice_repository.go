package repository

import (
	"context"
	"database/sql"
	"fmt"

	"learning_go/domain"
)

type invoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) domain.InvoiceRepository {
	return &invoiceRepository{
		db: db,
	}
}

func (ir *invoiceRepository) GetByOrderID(ctx context.Context, orderID int64) ([]*domain.Invoice, error) {
	query := fmt.Sprintf(`
		SELECT id, order_id, invoice_no, amount, status
		FROM %s
		WHERE order_id = ?
		ORDER BY id DESC
	`, domain.TableInvoice)

	rows, err := ir.db.QueryContext(ctx, query, orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to query invoices: %w", err)
	}
	defer rows.Close()

	var invoices []*domain.Invoice
	for rows.Next() {
		invoice := &domain.Invoice{}
		var invoiceNo sql.NullString
		var amount sql.NullFloat64
		var status sql.NullString

		if err := rows.Scan(
			&invoice.ID,
			&invoice.OrderID,
			&invoiceNo,
			&amount,
			&status,
		); err != nil {
			return nil, fmt.Errorf("failed to scan invoice: %w", err)
		}

		if invoiceNo.Valid {
			invoice.InvoiceNo = &invoiceNo.String
		}
		if amount.Valid {
			invoice.Amount = &amount.Float64
		}
		if status.Valid {
			invoice.Status = &status.String
		}

		invoices = append(invoices, invoice)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating invoices: %w", err)
	}

	return invoices, nil
}
