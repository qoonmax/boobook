package repository

import (
	"errors"
	"socialNetwork/internal/repository/model"
)

type UserRepository interface {
	Get(id int) (*model.User, error)
	Create(user *model.User) error
}

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)
