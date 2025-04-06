package service

import (
	"fmt"
	"gin-test/internal/dto"
	"gin-test/internal/model"
	"gin-test/internal/repository"
)

type PostService struct {
	PostRepo *repository.PostRepository
	UserRepo *repository.UserRepository
}

func NewPostService(postRepo *repository.PostRepository, userRepo *repository.UserRepository) *PostService {
	return &PostService{
		PostRepo: postRepo,
		UserRepo: userRepo,
	}
}

func (postService *PostService) CreatePost(email string, req *dto.CreatePostRequest) (*model.Post, error) {
	user, err := postService.UserRepo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("없는 이메일: %w", err)
	}

	post := &model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  user.ID,
	}

	if err := postService.PostRepo.Create(post); err != nil {
		return nil, fmt.Errorf("게시글 생성 실패: %w", err)
	}
	return post, nil

}

func (postService *PostService) GetAllPosts() ([]model.Post, error) {
	return postService.PostRepo.FindAll()
}

func (postService *PostService) GetPosts(email string) ([]model.Post, error) {
	user, err := postService.UserRepo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("없는 이메일: %w", err)
	}

	return postService.PostRepo.FindPostsByUserID(user.ID)
}
