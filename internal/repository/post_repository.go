package repository

import (
	"gin-test/internal/model"
	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (postRepository *PostRepository) Create(post *model.Post) error {
	return postRepository.DB.Create(post).Error
}

func (postRepository *PostRepository) FindAll() ([]model.Post, error) {
	var posts []model.Post
	if err := postRepository.DB.Preload("User").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (postRepository *PostRepository) FindPostsByUserID(userID uint) ([]model.Post, error) {
	var posts []model.Post
	if err := postRepository.DB.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
