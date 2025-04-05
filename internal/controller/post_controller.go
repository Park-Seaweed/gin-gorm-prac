package controller

import (
	"gin-test/internal/dto"
	"gin-test/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PostController struct {
	Service *service.PostService
}

func NewPostController(service *service.PostService) *PostController {
	return &PostController{Service: service}
}

func (postController *PostController) CreatePost(ctx *gin.Context) {
	userIDParam := ctx.Param("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 user_id입니다"})
		return
	}
	var req dto.CreatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 요청입니다", "details": err.Error()})
		return
	}

	req.UserID = uint(userID)

	post, err := postController.Service.CreatePost(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "게시글 생성 실패"})
		return
	}

	ctx.JSON(http.StatusCreated, post)

}

func (postController *PostController) GetPost(ctx *gin.Context) {
	//idParam := ctx.Param("id")
	//id, err := strconv.Atoi(idParam)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
	//	return
	//}
	//
	//for _, p := range model.Posts {
	//	if p.ID == id {
	//		ctx.JSON(http.StatusOK, p)
	//		return
	//	}
	//}
	//
	//
	//ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}

func (postController *PostController) GetAllPosts(c *gin.Context) {
	posts, err := postController.Service.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "게시글을 불러오지 못했습니다"})
		return
	}

	c.JSON(http.StatusOK, posts)
}
