package router

import (
	"gin-test/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, co *controller.Controllers) {

	api := r.Group("/api")

	RegisterUserRoutes(api, co.UserController)
	RegisterPostRoutes(api, co.PostController)

}
