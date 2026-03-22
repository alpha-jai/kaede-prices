package route

import (
	"github.com/koyo/kaede-prices/api/handler"
	"github.com/koyo/kaede-prices/repo"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	// Register watchlist routes
	watchlistHandler := handler.NewWatchlistHandler()
	RegisterWatchlistRoutes(router, watchlistHandler)

	// Register reward routes
	RegisterRewardRoutes(router)

	// Register auth routes
	userRepo := repo.NewUserRepository()
	authHandler := handler.NewAuthHandler(userRepo)
	RegisterAuthRoutes(router, authHandler)
}
