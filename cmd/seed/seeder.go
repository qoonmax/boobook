package main

import (
	"boobook/internal/config"
	"boobook/internal/repository/model"
	"boobook/internal/repository/postgres"
	"boobook/internal/slogger"
	"database/sql"
	"fmt"
	"github.com/go-faker/faker/v4"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

const (
	userCount = 1000
	batchSize = 5000
)

func main() {
	logger := slogger.NewLogger()

	cfg := config.MustLoad()

	dbWriteConnection, err := postgres.NewWriteConnection(cfg.WriteDBString)
	if err != nil {
		panic(fmt.Errorf("failed to connect to the database: %w", err))
	}

	defer func(dbWriteConnection *sql.DB) {
		if err = postgres.CloseConnection(dbWriteConnection); err != nil {
			logger.Error("failed to close the database connection", slogger.Err(err))
			return
		}
	}(dbWriteConnection)

	// Генерация и вставка пользователей
	fmt.Println("Starting to insert users...")
	err = generateAndInsertUsers(dbWriteConnection)
	if err != nil {
		panic(fmt.Errorf("failed to insert users: %w", err))
	}
	fmt.Println("Users have been successfully inserted.")

	// Генерация и вставка постов
	fmt.Println("Starting to insert posts...")
	err = generateAndInsertPosts(dbWriteConnection)
	if err != nil {
		panic(fmt.Errorf("failed to insert posts: %w", err))
	}
	fmt.Println("Posts have been successfully inserted.")
}

func generateAndInsertUsers(db *sql.DB) error {
	genders := []model.Gender{model.Male, model.Female, model.Other}
	cities := []string{"New York", "Los Angeles", "Chicago", "Houston", "Phoenix", "Philadelphia", "San Antonio", "San Diego", "Dallas", "San Jose"}

	for i := 0; i < userCount; i += batchSize {
		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("failed to start transaction: %w", err)
		}

		query := "INSERT INTO users (email, password, first_name, last_name, date_of_birth, gender, interests, city) VALUES "
		var values []interface{}
		var valueStrings []string

		for j := 0; j < batchSize && i+j < userCount; j++ {
			// Создание пользователя для тестирования
			if j == 0 {
				valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
					len(values)+1, len(values)+2, len(values)+3, len(values)+4, len(values)+5, len(values)+6, len(values)+7, len(values)+8))
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
				if err != nil {
					return err
				}
				values = append(values, "admin@admin.com", hashedPassword, "John", "Doe", time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), genders[1], "programming", cities[0])
				continue
			}

			dateOfBirth, err := time.Parse(time.DateOnly, faker.Date())
			if err != nil {
				return fmt.Errorf("failed to parse date of birth: %w", err)
			}

			randomGenderIdx := rand.Intn(len(genders))
			randomCityIdx := rand.Intn(len(cities))

			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
				len(values)+1, len(values)+2, len(values)+3, len(values)+4, len(values)+5, len(values)+6, len(values)+7, len(values)+8))
			values = append(values, faker.Email(), faker.Password(), strings.ToLower(faker.FirstName()), strings.ToLower(faker.LastName()), dateOfBirth, genders[randomGenderIdx], faker.Word(), cities[randomCityIdx])
		}

		query += strings.Join(valueStrings, ", ")

		_, err = tx.Exec(query, values...)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to execute batch insert: %w", err)
		}

		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("failed to commit transaction: %w", err)
		}
	}
	return nil
}

func generateAndInsertPosts(db *sql.DB) error {
	for i := 0; i < userCount*10; i += batchSize {
		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("failed to start transaction: %w", err)
		}

		query := "INSERT INTO posts (user_id, title, body) VALUES "
		var values []interface{}
		var valueStrings []string

		for j := 0; j < batchSize && i+j < userCount*10; j++ {
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)",
				len(values)+1, len(values)+2, len(values)+3))
			values = append(values, rand.Intn(userCount)+1, faker.Sentence(), faker.Paragraph())
		}

		query += strings.Join(valueStrings, ", ")

		_, err = tx.Exec(query, values...)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to execute batch insert: %w", err)
		}

		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("failed to commit transaction: %w", err)
		}
	}
	return nil
}
