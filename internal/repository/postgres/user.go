package postgres

import (
	"boobook/internal/repository"
	"boobook/internal/repository/model"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"strconv"
	"strings"
)

type userRepository struct {
	DBReadConnection  *sql.DB
	DBWriteConnection *sql.DB
}

func NewUserRepository(DBReadConnection *sql.DB, DBWriteConnection *sql.DB) repository.UserRepository {
	return &userRepository{
		DBReadConnection:  DBReadConnection,
		DBWriteConnection: DBWriteConnection,
	}
}

// TODO: Создать нормальные ошибки
func (r *userRepository) Create(user *model.User) error {
	const fnErr = "repository.postgres.userRepository.Create"

	var exists bool
	err := r.DBReadConnection.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", user.Email).Scan(&exists)
	if err != nil {
		return fmt.Errorf("(%s) error executing user existence check query: %w", fnErr, err)
	}

	if exists {
		return repository.ErrUserAlreadyExists
	}

	stmt, err := r.DBWriteConnection.Prepare("INSERT INTO users (email, password, first_name, last_name, date_of_birth, gender, interests, city) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
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

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	const fnErr = "repository.postgres.userRepository.GetByEmail"

	var user model.User
	query := `
		SELECT id, email, password, first_name, last_name, date_of_birth, gender, interests, city, created_at, updated_at 
		FROM users 
		WHERE email=$1
	`
	err := r.DBReadConnection.QueryRow(query, email).Scan(
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

func (r *userRepository) Get(id uint) (*model.User, error) {
	const fnErr = "repository.postgres.userRepository.Get"

	var user model.User
	query := `
		SELECT id, email, password, first_name, last_name, date_of_birth, gender, interests, city, created_at, updated_at 
		FROM users 
		WHERE id=$1
	`
	err := r.DBReadConnection.QueryRow(query, id).Scan(
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

func (r *userRepository) Search(firstName, lastName string) ([]*model.User, error) {
	const fnErr = "repository.postgres.userRepository.Search"

	var users []*model.User

	query := `
		SELECT id, email, password, first_name, last_name, date_of_birth, gender, interests, city, created_at, updated_at 
		FROM users
	`

	var rows *sql.Rows
	var err error
	var filters []string
	var args []interface{}

	if lastName != "" {
		lastName = strings.ToLower(lastName)
		filters = append(filters, "last_name LIKE $"+strconv.Itoa(len(filters)+1))
		args = append(args, lastName+"%")
	}

	if firstName != "" {
		firstName = strings.ToLower(firstName)
		filters = append(filters, "first_name LIKE $"+strconv.Itoa(len(filters)+1))
		args = append(args, firstName+"%")
	}

	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}

	fmt.Println(query, args)
	rows, err = r.DBReadConnection.Query(query, args...)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", fnErr, err)
	}

	defer func() {
		if err = rows.Close(); err != nil {
			err = fmt.Errorf("%s: %w", fnErr, err)
		}
	}()

	for rows.Next() {
		var user model.User
		if err = rows.Scan(
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
		); err != nil {
			return nil, fmt.Errorf("%s: %w", fnErr, err)
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", fnErr, err)
	}

	return users, err
}
