package service

import (
	"fmt"
	"gin-test/internal/dto"
	"gin-test/internal/model"
	"gin-test/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo   *repository.UserRepository
	JwtService *JwtService
}

func NewUserService(repo *repository.UserRepository, jwtService *JwtService) *UserService {
	return &UserService{
		UserRepo:   repo,
		JwtService: jwtService,
	}
}

func (userService *UserService) RegisterUser(req *dto.CreateUserRequest) (*dto.UserTokenResponse, error) {
	existingUser, err := userService.UserRepo.FindByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("이미 가입된 이메일입니다")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("비밀번호 해싱 실패: %w", err)
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashed),
	}

	if err := userService.UserRepo.Create(user); err != nil {
		return nil, fmt.Errorf("사용자 생성 실패: %w", err)
	}

	accessToken, err := userService.JwtService.GenerateAccessToken(user.Email)
	if err != nil {
		return nil, fmt.Errorf("access 토큰 생성 실패: %w", err)
	}
	refreshTokenStr, err := userService.JwtService.GenerateRefreshToken(user.Email)
	if err != nil {
		return nil, fmt.Errorf("refresh 토큰 생성 실패: %w", err)
	}

	if err := userService.UserRepo.UpdateRefreshToken(user.ID, refreshTokenStr); err != nil {
		return nil, fmt.Errorf("refresh 토큰 저장 실패: %w", err)
	}

	return &dto.UserTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenStr,
	}, nil
}

func (userService *UserService) GetUserByID(id uint) (*model.User, error) {
	return userService.UserRepo.FindByID(id)
}

func (userService *UserService) LoginUser(req *dto.LoginUserRequest) (*dto.UserTokenResponse, error) {
	user, err := userService.UserRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("이메일이 존재하지 않습니다: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("비밀번호가 틀렸습니다: %w", err)
	}

	accessToken, err := userService.JwtService.GenerateAccessToken(user.Email)
	if err != nil {
		return nil, fmt.Errorf("access 토큰 생성 실패: %w", err)
	}
	refreshToken, err := userService.JwtService.GenerateRefreshToken(user.Email)
	if err != nil {
		return nil, fmt.Errorf("refresh 토큰 생성 실패: %w", err)
	}

	if err := userService.UserRepo.UpdateRefreshToken(user.ID, refreshToken); err != nil {
		return nil, fmt.Errorf("refresh 토큰 저장 실패: %w", err)
	}

	return &dto.UserTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
