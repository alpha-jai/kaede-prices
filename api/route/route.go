package route

import (
	"github.com/koyo/kaede-prices/api/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	// Register watchlist routes
	watchlistHandler := handler.NewWatchlistHandler()
	RegisterWatchlistRoutes(router, watchlistHandler)
}
