package model

import "time"

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = "other"
)

type User struct {
	ID          int       `json:"id"`
	Email       string    `json:"email" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	FirstName   string    `json:"first_name" binding:"required"`
	LastName    string    `json:"last_name" binding:"required"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	Gender      Gender    `json:"gender" binding:"required"`
	Interests   string    `json:"interests" binding:"required"`
	City        string    `json:"city" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
