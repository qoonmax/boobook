package service

import (
	"boobook/internal/http/request"
	"boobook/internal/repository"
	"boobook/internal/repository/model"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (s *authService) Register(registerRequest *request.RegisterRequest) error {
	const fnErr = "service.userService.Register"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Email:       registerRequest.Email,
		Password:    string(hashedPassword),
		FirstName:   registerRequest.FirstName,
		LastName:    registerRequest.LastName,
		DateOfBirth: registerRequest.DateOfBirth,
		Gender:      model.Gender(registerRequest.Gender),
		Interests:   registerRequest.Interests,
		City:        registerRequest.City,
	}

	return s.userRepository.Create(user)
}

func (s *authService) Login(loginRequest *request.LoginRequest) (string, error) {
	const fnErr = "service.userService.Login"

	user, err := s.userRepository.GetByEmail(loginRequest.Email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return "", ErrInvalidPassword
	}

	payload := jwt.MapClaims{
		"user_email": user.Email,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// TODO: заменить на переменную окружения
	encryptedToken, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "", fmt.Errorf("(%s) error signing token: %w", fnErr, err)
	}

	return encryptedToken, nil
}
