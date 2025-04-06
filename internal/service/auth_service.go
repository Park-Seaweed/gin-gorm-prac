package service

import (
	"fmt"
	"gin-test/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	*config.Config
}

func NewJwtServiceFromEnv(cfg *config.Config) *JwtService {
	return &JwtService{cfg}
}

func (jwtService *JwtService) GenerateAccessToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email":      email,
		"token_type": "access",
		"exp":        time.Now().Add(60 * time.Minute).Unix(),
		"iat":        time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtService.JwtSecret.Secret))
}

func (jwtService *JwtService) GenerateRefreshToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email":      email,
		"token_type": "refresh",
		"exp":        time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat":        time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtService.JwtSecret.Secret))
}

func (jwtService *JwtService) ParseToken(tokenStr string) (email string, tokenType string, err error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		return []byte(jwtService.JwtSecret.Secret), nil
	})
	if err != nil || !token.Valid {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", err
	}

	email, _ = claims["email"].(string)
	tokenType, _ = claims["token_type"].(string)
	return email, tokenType, nil
}

func (jwtService *JwtService) RefreshToken(refreshToken string, tokenKind string) (string, error) {
	email, tokenType, err := jwtService.ParseToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("토큰 파싱 실패: %w", err)
	}
	if tokenType != "refresh" {
		return "", fmt.Errorf("리프레시 토큰이 아님")
	}

	switch tokenKind {
	case "access":
		return jwtService.GenerateAccessToken(email)
	case "refresh":
		return jwtService.GenerateRefreshToken(email)
	default:
		return "", fmt.Errorf("알 수 없는 토큰 종류입니다")
	}
}
