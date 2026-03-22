package route

import (
	"github.com/gin-gonic/gin"
	"github.com/koyo/kaede-prices/api/handler"
)

func RegisterRewardRoutes(r *gin.RouterGroup, h *handler.RewardHandler) {
	rewardGroup := r.Group("/api/v1/rewards")
	{
		rewardGroup.GET("/", h.GetRewards)
		rewardGroup.POST("/claim", h.ClaimReward)
	}
}
