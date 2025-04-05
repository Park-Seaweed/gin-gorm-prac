package router

import (
	"gin-test/internal/controller"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, uc *controller.UserController) {
	user := rg.Group("/users")
	{
		user.POST("/register", uc.CreateUser)
		user.GET("/:id", uc.GetUser)
		user.POST("/login", uc.LoginUser)
	}
}
