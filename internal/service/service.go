package service

import "socialNetwork/internal/repository/model"

type UserService interface {
	Get(id int) (*model.User, error)
	Create(user *model.User) error
}
