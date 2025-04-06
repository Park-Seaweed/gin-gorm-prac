package service

import (
	"gin-test/internal/config"
	"testing"
)

func TestJwtService_WithoutTestify(t *testing.T) {
	cfg := &config.Config{}
	cfg.JwtSecret.Secret = "test-secret-key"
	jwtService := NewJwtServiceFromEnv(cfg)
	email := "hello@example.com"

	accessToken, err := jwtService.GenerateAccessToken(email)
	if err != nil {
		t.Errorf("access token 생성 실패: %v", err)
	}
	if accessToken == "" {
		t.Errorf("access token 비어있음")
	}

	parsedEmail, tokenType, err := jwtService.ParseToken(accessToken)
	if err != nil {
		t.Errorf("access token 파싱 실패: %v", err)
	}
	if parsedEmail != email {
		t.Errorf("email 불일치: got %s, want %s", parsedEmail, email)
	}
	if tokenType != "access" {
		t.Errorf("tokenType 불일치: got %s, want access", tokenType)
	}
}
