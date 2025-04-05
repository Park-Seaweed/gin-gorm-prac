package service

import (
	"fmt"
	"gin-test/internal/dto"
	"gin-test/internal/model"
	"gin-test/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (userService *UserService) RegisterUser(req *dto.CreateUserRequest) (*model.User, error) {
	existingUser, err := userService.Repo.FindByEmail(req.Email)
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

	if err := userService.Repo.Create(user); err != nil {
		return nil, fmt.Errorf("사용자 생성 실패: %w", err)
	}

	return user, nil
}

func (userService *UserService) GetUserByID(id uint) (*model.User, error) {
	return userService.Repo.FindByID(id)
}

func (userService *UserService) LoginUser(req *dto.LoginUserRequest) (*model.User, error) {
	user, err := userService.Repo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	return user, nil
}
