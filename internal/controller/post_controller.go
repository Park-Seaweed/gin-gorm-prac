package controller

import (
	"gin-test/internal/dto"
	"gin-test/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostController struct {
	Service *service.PostService
}

func NewPostController(service *service.PostService) *PostController {
	return &PostController{Service: service}
}

func (postController *PostController) CreatePost(ctx *gin.Context) {
	emailValue, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "로그인 정보가 없습니다"})
		return
	}

	email := emailValue.(string)

	var req dto.CreatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 요청입니다", "details": err.Error()})
		return
	}

	post, err := postController.Service.CreatePost(email, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "게시글 생성 실패"})
		return
	}

	ctx.JSON(http.StatusCreated, post)

}

func (postController *PostController) GetPost(ctx *gin.Context) {
	emailValue, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "로그인 정보가 없습니다"})
		return
	}

	email := emailValue.(string)

	userPosts, err := postController.Service.GetPosts(email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, userPosts)
}

func (postController *PostController) GetAllPosts(c *gin.Context) {
	posts, err := postController.Service.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "게시글을 불러오지 못했습니다"})
		return
	}

	c.JSON(http.StatusOK, posts)
}
