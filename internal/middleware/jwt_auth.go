package middleware

import (
	"gin-test/internal/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(jwtService *service.JwtService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization 헤더가 필요합니다"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		email, tokenType, err := jwtService.ParseToken(tokenStr)
		if err != nil || tokenType != "access" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "유효하지 않은 액세스 토큰입니다"})
			return
		}

		// 유저 식별 정보 컨텍스트에 저장
		ctx.Set("email", email)
		ctx.Next()
	}
}
