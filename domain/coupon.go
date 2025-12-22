package domain

import "context"

const (
	TableCoupon = "oms_coupons"
)

// Coupon represents a user's coupon detail record.
// Adjust fields according to the actual oms_coupons table schema if needed.
type Coupon struct {
	ID             int64    `json:"id" db:"id"`
	UserID         int64    `json:"user_id" db:"user_id"`
	CouponCode     *string  `json:"coupon_code,omitempty" db:"coupon_code"`
	DiscountAmount *float64 `json:"discount_amount,omitempty" db:"discount_amount"`
	Status         *string  `json:"status,omitempty" db:"status"`
}

// UserCouponRequest is the request body for querying coupons by user.
type UserCouponRequest struct {
	UserID int64 `json:"user_id" binding:"required"`
}

type CouponRepository interface {
	// GetByUserID returns all coupons for a given user.
	GetByUserID(ctx context.Context, userID int64) ([]*Coupon, error)
}

type CouponUsecase interface {
	// GetByUserID returns all coupons for a given user.
	GetByUserID(ctx context.Context, userID int64) ([]*Coupon, error)
}
