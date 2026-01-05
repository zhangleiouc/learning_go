package usecase

import (
	"context"
	"time"

	"learning_go/domain"
)

type cartUsecase struct {
	cartRepository domain.CartRepository
	contextTimeout time.Duration
}

func NewCartUsecase(cartRepository domain.CartRepository, timeout time.Duration) domain.CartUsecase {
	return &cartUsecase{
		cartRepository: cartRepository,
		contextTimeout: timeout,
	}
}

func (cu *cartUsecase) AddOrUpdate(c context.Context, userID, productID, quantity int64) (int64, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.cartRepository.AddOrUpdate(ctx, userID, productID, quantity)
}

func (cu *cartUsecase) GetByUserID(c context.Context, userID int64) ([]*domain.CartItem, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.cartRepository.GetByUserID(ctx, userID)
}

func (cu *cartUsecase) Delete(c context.Context, cartID int64) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.cartRepository.Delete(ctx, cartID)
}
