package repo

import "github.com/koyo/kaede-prices/domain"

type RewardRepo interface {
	CreateTransaction(t *domain.Transaction) error
	GetTotalPoints(userID uint) (int, error)
	AddReward(r *domain.UserReward) error
}
