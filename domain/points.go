package domain

import "context"

const (
	TableUserPoints        = "user_points"
	TableUserPointsHistory = "user_points_history"

	// 积分类型
	PointsTypeEarned   = "EARNED"   // 获得积分
	PointsTypeConsumed = "CONSUMED" // 消费积分
)

// UserPoints represents user's points balance.
type UserPoints struct {
	ID        int64   `json:"id" db:"id"`
	UserID    int64   `json:"user_id" db:"user_id"`
	Points    int64   `json:"points" db:"points"`
	UpdatedAt *string `json:"updated_at,omitempty" db:"updated_at"`
}

// PointsHistory represents a points transaction record.
type PointsHistory struct {
	ID          int64   `json:"id" db:"id"`
	UserID      int64   `json:"user_id" db:"user_id"`
	Points      int64   `json:"points" db:"points"`
	OrderID     *int64  `json:"order_id,omitempty" db:"order_id"`
	Type        string  `json:"type" db:"type"`
	Description *string `json:"description,omitempty" db:"description"`
	CreatedAt   *string `json:"created_at,omitempty" db:"created_at"`
}

type PointsRepository interface {
	// AddPoints adds points to user's balance and records history.
	AddPoints(ctx context.Context, userID int64, points int64, orderID *int64, description string) error
	// GetByUserID returns user's points balance.
	GetByUserID(ctx context.Context, userID int64) (*UserPoints, error)
	// GetHistoryByUserID returns user's points history records.
	GetHistoryByUserID(ctx context.Context, userID int64) ([]*PointsHistory, error)
}

type PointsUsecase interface {
	// AddPoints adds points to user's balance and records history.
	AddPoints(ctx context.Context, userID int64, points int64, orderID *int64, description string) error
	// GetByUserID returns user's points balance.
	GetByUserID(ctx context.Context, userID int64) (*UserPoints, error)
	// GetHistoryByUserID returns user's points history records.
	GetHistoryByUserID(ctx context.Context, userID int64) ([]*PointsHistory, error)
}
