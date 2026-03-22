package repo

import (
	"context"
	"kaede-prices/domain"
)

type RewardRepository interface {
	GetReward(ctx context.Context, userID string) (*domain.UserReward, error)
	UpdatePoints(ctx context.Context, userID string, amount int) error
	AddTransaction(ctx context.Context, tx domain.Transaction) error
}

type rewardRepository struct {
	// db *sql.DB
}

func NewRewardRepository() RewardRepository {
	return &rewardRepository{}
}

func (r *rewardRepository) GetReward(ctx context.Context, userID string) (*domain.UserReward, error) {
	return &domain.UserReward{UserID: userID, Points: 0}, nil
}

func (r *rewardRepository) UpdatePoints(ctx context.Context, userID string, amount int) error {
	return nil
}

func (r *rewardRepository) AddTransaction(ctx context.Context, tx domain.Transaction) error {
	return nil
}
