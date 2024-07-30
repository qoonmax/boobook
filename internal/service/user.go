package service

import (
	"socialNetwork/internal/repository"
	"socialNetwork/internal/repository/model"
)

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		UserRepository: userRepository,
	}
}

func (s *userService) Get(id int) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) Create(user *model.User) error {
	return s.UserRepository.Create(user)
}
