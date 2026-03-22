package domain

import "time"

type UserReward struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Points    int       `json:"points"`
	CreatedAt time.Time `json:"created_at"`
}

type Transaction struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Action    string    `json:"action"` // e.g., "report_price"
	Points    int       `json:"points"`
	CreatedAt time.Time `json:"created_at"`
}
