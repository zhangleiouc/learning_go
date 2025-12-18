package usecase

import (
	"context"
	"time"

	"learning_go/domain"
)

type invoiceUsecase struct {
	invoiceRepository domain.InvoiceRepository
	contextTimeout    time.Duration
}

func NewInvoiceUsecase(invoiceRepository domain.InvoiceRepository, timeout time.Duration) domain.InvoiceUsecase {
	return &invoiceUsecase{
		invoiceRepository: invoiceRepository,
		contextTimeout:    timeout,
	}
}

func (iu *invoiceUsecase) GetByOrderID(c context.Context, orderID int64) ([]*domain.Invoice, error) {
	ctx, cancel := context.WithTimeout(c, iu.contextTimeout)
	defer cancel()
	return iu.invoiceRepository.GetByOrderID(ctx, orderID)
}
