package router

import (
	"gin-test/internal/controller"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, ac *controller.AuthController) {
	auth := rg.Group("/auth/token")
	{
		auth.POST("/access", ac.RefreshAccessToken)
		auth.POST("/refresh", ac.RefreshRefreshToken)
	}
}
