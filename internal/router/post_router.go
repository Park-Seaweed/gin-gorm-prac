package router

import (
	"gin-test/internal/controller"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(rg *gin.RouterGroup, pc *controller.PostController) {
	users := rg.Group("/users/:user_id")
	{
		users.POST("/posts", pc.CreatePost)
		//users.GET("/posts", pc.GetUserPosts)
	}

	posts := rg.Group("/posts")
	{
		posts.GET("/:id", pc.GetPost)
		posts.GET("/", pc.GetAllPosts)
	}
}
