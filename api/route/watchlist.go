package route

import (
	"github.com/koyo/kaede-prices/api/handler"
	"github.com/gin-gonic/gin"
)

func RegisterWatchlistRoutes(router *gin.RouterGroup, handler *handler.WatchlistHandler) {
	watchlist := router.Group("/watchlists")
	{
		watchlist.POST("", handler.Create)
		watchlist.GET("", handler.List)
		watchlist.DELETE("/:id", handler.Delete)
	}
}
