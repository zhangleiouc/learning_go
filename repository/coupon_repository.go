package repository

import (
	"context"
	"database/sql"
	"fmt"

	"learning_go/domain"
)

type couponRepository struct {
	db *sql.DB
}

func NewCouponRepository(db *sql.DB) domain.CouponRepository {
	return &couponRepository{
		db: db,
	}
}

func (cr *couponRepository) GetByUserID(ctx context.Context, userID int64) ([]*domain.Coupon, error) {
	query := fmt.Sprintf(`
		SELECT id, user_id, coupon_code, discount_amount, status
		FROM %s
		WHERE user_id = ?
		ORDER BY id DESC
	`, domain.TableCoupon)

	rows, err := cr.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query coupons: %w", err)
	}
	defer rows.Close()

	var coupons []*domain.Coupon
	for rows.Next() {
		coupon := &domain.Coupon{}
		var couponCode sql.NullString
		var discountAmount sql.NullFloat64
		var status sql.NullString

		if err := rows.Scan(
			&coupon.ID,
			&coupon.UserID,
			&couponCode,
			&discountAmount,
			&status,
		); err != nil {
			return nil, fmt.Errorf("failed to scan coupon: %w", err)
		}

		if couponCode.Valid {
			coupon.CouponCode = &couponCode.String
		}
		if discountAmount.Valid {
			coupon.DiscountAmount = &discountAmount.Float64
		}
		if status.Valid {
			coupon.Status = &status.String
		}

		coupons = append(coupons, coupon)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating coupons: %w", err)
	}

	return coupons, nil
}
