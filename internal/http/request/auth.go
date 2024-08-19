package request

import "time"

type RegisterRequest struct {
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password" binding:"required,min=8"`
	FirstName   string    `json:"first_name" binding:"required"`
	LastName    string    `json:"last_name" binding:"required"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	Gender      string    `json:"gender" binding:"required,oneof=male female other"`
	Interests   string    `json:"interests" binding:"required"`
	City        string    `json:"city" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
