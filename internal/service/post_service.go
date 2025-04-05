package service

import (
	"fmt"
	"gin-test/internal/dto"
	"gin-test/internal/model"
	"gin-test/internal/repository"
)

type PostService struct {
	Repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{Repo: repo}
}

func (postService *PostService) CreatePost(req *dto.CreatePostRequest) (*model.Post, error) {
	post := &model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}

	if err := postService.Repo.Create(post); err != nil {
		return nil, fmt.Errorf("게시글 생성 실패: %w", err)
	}
	return post, nil

}

func (postService *PostService) GetAllPosts() ([]model.Post, error) {
	return postService.Repo.FindAll()
}
