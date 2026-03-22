package service

import (
	"github.com/koyo/kaede-prices/domain"
	"github.com/koyo/kaede-prices/repo"
)

type RewardService struct {
	repo repo.RewardRepo
}

func NewRewardService(r repo.RewardRepo) *RewardService {
	return &RewardService{repo: r}
}

func (s *RewardService) ReportPrice(userID uint) error {
	t := &domain.Transaction{
		UserID: userID,
		Action: "report_price",
		Points: 10,
	}
	return s.repo.CreateTransaction(t)
}
