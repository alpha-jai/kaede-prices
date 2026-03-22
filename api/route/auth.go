package route

import (
	"github.com/koyo/kaede-prices/api/handler"
	"github.com/koyo/kaede-prices/domain"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup, handler *handler.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}
}
