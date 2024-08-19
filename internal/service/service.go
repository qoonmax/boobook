package service

import (
	"boobook/internal/http/request"
	"boobook/internal/repository/model"
	"errors"
)

var (
	ErrInvalidPassword = errors.New("invalid password")
)

type AuthService interface {
	Login(loginRequest *request.LoginRequest) (string, error)
	Register(registerRequest *request.RegisterRequest) error
}

type UserService interface {
	Get(id uint) (*model.User, error)
}
