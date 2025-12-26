package usecase

import (
	"context"
	"time"

	"learning_go/domain"
)

type couponUsecase struct {
	couponRepository domain.CouponRepository
	contextTimeout   time.Duration
}

func NewCouponUsecase(couponRepository domain.CouponRepository, timeout time.Duration) domain.CouponUsecase {
	return &couponUsecase{
		couponRepository: couponRepository,
		contextTimeout:   timeout,
	}
}

func (cu *couponUsecase) GetByUserID(c context.Context, userID int64) ([]*domain.Coupon, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.couponRepository.GetByUserID(ctx, userID)
}
