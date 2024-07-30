package postgres

import (
	"database/sql"
	"fmt"
	"socialNetwork/internal/repository"
	"socialNetwork/internal/repository/model"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db: db}
}

// TODO: Создать нормальные ошибки
func (r *userRepository) Create(user *model.User) error {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", user.Email).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to query row: %v", err)
	}

	if exists {
		return fmt.Errorf("user already exists")
	}

	stmt, err := r.db.Prepare("INSERT INTO users (email, password, first_name, last_name, date_of_birth, gender, interests, city) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}

	_, err = stmt.Exec(user.Email, user.Password, user.FirstName, user.LastName, user.DateOfBirth, user.Gender, user.Interests, user.City)
	//TODO: Check if the user already exists, return ErrUserAlreadyExists
	if err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}

	return nil
}

func (r *userRepository) Get(id int) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
