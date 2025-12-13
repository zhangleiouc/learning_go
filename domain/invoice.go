package domain

import "context"

const (
	TableInvoice = "oms_invoices"
)

// Invoice represents invoice data tied to an order.
// If your table has additional columns, extend this struct and repository scan accordingly.
type Invoice struct {
	ID        int64    `json:"id" db:"id"`
	OrderID   int64    `json:"order_id" db:"order_id"`
	InvoiceNo *string  `json:"invoice_no,omitempty" db:"invoice_no"`
	Amount    *float64 `json:"amount,omitempty" db:"amount"`
	Status    *string  `json:"status,omitempty" db:"status"`
}

type InvoiceRepository interface {
	GetByOrderID(ctx context.Context, orderID int64) ([]*Invoice, error)
}

type InvoiceUsecase interface {
	GetByOrderID(ctx context.Context, orderID int64) ([]*Invoice, error)
}
