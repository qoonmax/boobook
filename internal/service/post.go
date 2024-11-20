package service

import (
	"boobook/internal/repository"
	"boobook/internal/repository/model"
	"fmt"
)

type postService struct {
	postRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) PostService {
	return &postService{
		postRepository: postRepository,
	}
}

func (s *postService) GetList() ([]*model.Post, error) {
	const fnErr = "service.postService.GetList"

	posts, err := s.postRepository.GetList()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fnErr, err)
	}

	return posts, nil
}
