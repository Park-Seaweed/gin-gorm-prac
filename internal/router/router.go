package router

import (
	"gin-test/internal/controller"
	"gin-test/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, co *controller.Controllers, jwtService *service.JwtService) {
	api := r.Group("/api")

	RegisterUserRoutes(api, co.UserController)
	RegisterPostRoutes(api, co.PostController, jwtService)
	RegisterAuthRoutes(api, co.AuthController)
}
