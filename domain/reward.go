package domain

import "time"

type TransactionType string

const (
	Earned   TransactionType = "EARNED"
	Redeemed TransactionType = "REDEEMED"
)

type Transaction struct {
	ID        string          `json:"id"`
	UserID    string          `json:"user_id"`
	Amount    int             `json:"amount"`
	Type      TransactionType `json:"type"`
	CreatedAt time.Time       `json:"created_at"`
}

type UserReward struct {
	UserID            string        `json:"user_id"`
	Points            int           `json:"points"`
	TransactionHistory []Transaction `json:"transaction_history"`
}
