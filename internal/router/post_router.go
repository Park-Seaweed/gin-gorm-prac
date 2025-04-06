package router

import (
	"gin-test/internal/controller"
	"gin-test/internal/middleware"
	"gin-test/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(rg *gin.RouterGroup, pc *controller.PostController, jwtService *service.JwtService) {
	users := rg.Group("/post")
	users.Use(middleware.JWTAuthMiddleware(jwtService))
	{
		users.POST("/", pc.CreatePost)
		users.GET("/", pc.GetPost)
	}

	posts := rg.Group("/posts")
	{

		posts.GET("/", pc.GetAllPosts)
	}
}
