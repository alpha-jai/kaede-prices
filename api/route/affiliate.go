package route

import (
	"github.com/gin-gonic/gin"
	"github.com/koyo/kaede-prices/api/handler"
)

func RegisterAffiliateRoutes(r *gin.RouterGroup, h *handler.AffiliateHandler) {
	affiliateGroup := r.Group("/api/v1/affiliate")
	{
		affiliateGroup.GET("/", h.GetLinks)
		affiliateGroup.GET("/track", h.TrackClick)
	}
}
