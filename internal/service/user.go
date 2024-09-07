package service

import (
	"boobook/internal/repository"
	"boobook/internal/repository/model"
	"fmt"
	"unicode"
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

func (s *userService) Search(firstName string, lastName string) ([]*model.User, error) {
	const fnErr = "service.userService.Search"

	if firstName != "" {
		runes := []rune(firstName)
		runes[0] = unicode.ToUpper(runes[0])
		firstName = string(runes)
	}

	if lastName != "" {
		runes := []rune(lastName)
		runes[0] = unicode.ToUpper(runes[0])
		lastName = string(runes)
	}

	users, err := s.userRepository.Search(firstName, lastName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fnErr, err)
	}

	return users, nil
}
