package domain

type Watchlist struct {
	ID             int64   `json:"id" db:"id" dbfield:"id"`
	UserID         int64   `json:"user_id" db:"user_id" dbfield:"user_id"`
	ItemID         int64   `json:"item_id" db:"item_id" dbfield:"item_id"`
	AlertThreshold float64 `json:"alert_threshold" db:"alert_threshold" dbfield:"alert_threshold"`
	CreatedAt      int64   `json:"created_at" db:"created_at" dbfield:"created_at"`
	UpdatedAt      int64   `json:"updated_at" db:"updated_at" dbfield:"updated_at"`
}
