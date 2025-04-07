package controller

import (
	"gin-test/internal/dto"
	"gin-test/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	JwtService *service.JwtService
}

func NewAuthController(jwtService *service.JwtService) *AuthController {
	return &AuthController{JwtService: jwtService}
}

func (authController *AuthController) RefreshAccessToken(ctx *gin.Context) {
	var req dto.TokenRefreshRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "refresh_token이 필요합니다"})
		return
	}

	newAccessToken, err := authController.JwtService.RefreshToken(req.RefreshToken, "access")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "토큰 재발급 실패", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}

//func (authController *AuthController) RefreshRefreshToken(ctx *gin.Context) {
//	var req dto.TokenRefreshRequest
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "refresh_token이 필요합니다"})
//		return
//	}
//
//	newRefreshToken, err := authController.JwtService.RefreshToken(req.RefreshToken, "refresh")
//	if err != nil {
//		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "리프레시 토큰 재발급 실패", "details": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{"refresh_token": newRefreshToken})
//}
