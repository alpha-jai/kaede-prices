package route

import (
	"github.com/koyo/kaede-prices/api/handler"
	"github.com/koyo/kaede-prices/repo"
	"github.com/koyo/kaede-prices/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	// Register watchlist routes
	watchlistHandler := handler.NewWatchlistHandler()
	RegisterWatchlistRoutes(router, watchlistHandler)

	// Register reward routes
	rewardSvc := service.GetRewardService()
	rewardHandler := handler.NewRewardHandler(rewardSvc)
	RegisterRewardRoutes(router, rewardHandler)

	// Register affiliate routes
	affiliateSvc := service.GetAffiliateService()
	affiliateHandler := handler.NewAffiliateHandler(affiliateSvc)
	RegisterAffiliateRoutes(router, affiliateHandler)

	// Register auth routes
	userRepo := repo.NewUserRepository()
	authHandler := handler.NewAuthHandler(userRepo)
	RegisterAuthRoutes(router, authHandler)
}
