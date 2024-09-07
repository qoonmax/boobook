package repository

import (
	"boobook/internal/repository/model"
	"errors"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)

type UserRepository interface {
	Get(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	Search(firstName, lastName string) ([]*model.User, error)
}
