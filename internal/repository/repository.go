package repository

import "socialNetwork/internal/repository/model"

type UserRepository interface {
	Get(id int) (*model.User, error)
	Create(user *model.User) error
}
