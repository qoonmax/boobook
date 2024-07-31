package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
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
	const fnErr = "repository.postgres.userRepository.Create"

	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", user.Email).Scan(&exists)
	if err != nil {
		return fmt.Errorf("(%s) error executing user existence check query: %w", fnErr, err)
	}

	if exists {
		return repository.ErrUserAlreadyExists
	}

	stmt, err := r.db.Prepare("INSERT INTO users (email, password, first_name, last_name, date_of_birth, gender, interests, city) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return fmt.Errorf("(%s) error preparing insert statement: %w", fnErr, err)
	}

	_, err = stmt.Exec(user.Email, user.Password, user.FirstName, user.LastName, user.DateOfBirth, user.Gender, user.Interests, user.City)
	if err != nil {
		var pq *pq.Error
		if errors.As(err, &pq) {
			if pq.Code == "23505" {
				return repository.ErrUserAlreadyExists
			}
		}

		return fmt.Errorf("(%s) error executing user creation query: %w", fnErr, err)
	}

	return nil
}

func (r *userRepository) Get(id int) (*model.User, error) {
	const fnErr = "repository.postgres.userRepository.Get"

	var user model.User
	query := `
		SELECT id, email, password, first_name, last_name, date_of_birth, gender, interests, city, created_at, updated_at 
		FROM users 
		WHERE id=$1
	`
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.DateOfBirth,
		&user.Gender,
		&user.Interests,
		&user.City,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrUserNotFound
		}

		return nil, fmt.Errorf("(%s) error getting user: %w", fnErr, err)
	}

	return &user, nil
}
