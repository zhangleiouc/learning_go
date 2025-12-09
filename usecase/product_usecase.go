package usecase

import (
	"context"
	"time"

	"learning_go/domain"
)

type productUsecase struct {
	productRepository domain.ProductRepository
	contextTimeout    time.Duration
}

func NewProductUsecase(productRepository domain.ProductRepository, timeout time.Duration) domain.ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
		contextTimeout:    timeout,
	}
}

func (pu *productUsecase) SearchByName(c context.Context, keyword string) ([]*domain.Product, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.productRepository.SearchByName(ctx, keyword)
}
