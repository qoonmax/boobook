package service

import (
	"boobook/internal/repository"
	"boobook/internal/repository/model"
	"fmt"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Get(id uint) (*model.User, error) {
	const fnErr = "service.userService.Get"

	user, err := s.userRepository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fnErr, err)
	}

	return user, nil
}
