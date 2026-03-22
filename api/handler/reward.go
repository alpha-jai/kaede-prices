package handler

import (
	"encoding/json"
	"net/http"
	"kaede-prices/api/payload"
	"kaede-prices/repo"
)

type RewardHandler struct {
	repo repo.RewardRepository
}

func NewRewardHandler(r repo.RewardRepository) *RewardHandler {
	return &RewardHandler{repo: r}
}

func (h *RewardHandler) GetRewards(w http.ResponseWriter, r *http.Request) {
	// mock
	json.NewEncoder(w).Encode(payload.RewardsListResponse{UserID: "1", Points: 100})
}

func (h *RewardHandler) ClaimReward(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
