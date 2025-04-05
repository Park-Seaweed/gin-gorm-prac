package repository

import (
	"fmt"
	"gin-test/internal/logger"
	"gin-test/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	logger.Log.Info("유저 조회 성공", zap.Any("user", user))
	return &user, nil
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("err: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) name() {

}
