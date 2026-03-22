package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koyo/kaede-prices/service"
)

type RewardHandler struct {
	svc *service.RewardService
}

func NewRewardHandler(svc *service.RewardService) *RewardHandler {
	return &RewardHandler{svc: svc}
}

func (h *RewardHandler) GetRewards(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
}

func (h *RewardHandler) ClaimReward(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Not implemented"})
}
