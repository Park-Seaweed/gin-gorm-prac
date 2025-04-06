package controller

import (
	"gin-test/internal/dto"
	"gin-test/internal/logger"
	"gin-test/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type UserController struct {
	Service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{Service: service}
}

func (userController *UserController) CreateUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "유효하지 않은 입력값입니다",
			"details": err.Error(),
		})
		return
	}

	res, err := userController.Service.RegisterUser(&req)
	if err != nil {
		logger.Log.Error("에러", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (userController *UserController) GetUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	user, err := userController.Service.GetUserByID(uint(id))
	if err != nil {
		logger.Log.Error("요청 파싱 실패", zap.Error(err))
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (userController *UserController) LoginUser(ctx *gin.Context) {
	var req dto.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "유효하지 않은 입력값입니다.",
			"details": err.Error(),
		})
		return
	}

	user, err := userController.Service.LoginUser(&req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "이메일 또는 비밀번호가 잘못되었습니다.",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
