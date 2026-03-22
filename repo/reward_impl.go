package repo

import (
	"github.com/koyo/kaede-prices/domain"
)

type rewardRepo struct {}

func NewRewardRepo() RewardRepo {
	return &rewardRepo{}
}

func (r *rewardRepo) CreateTransaction(t *domain.Transaction) error {
	return nil
}
func (r *rewardRepo) GetTotalPoints(userID uint) (int, error) {
	return 0, nil
}
func (r *rewardRepo) AddReward(rd *domain.UserReward) error {
	return nil
}
