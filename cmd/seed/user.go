package main

import (
	"boobook/internal/config"
	"boobook/internal/repository/model"
	"boobook/internal/repository/postgres"
	"boobook/internal/slogger"
	"database/sql"
	"fmt"
	"github.com/go-faker/faker/v4"
	"math/rand"
	"strings"
	"time"
)

const (
	userCount = 1000000
	batchSize = 5000
)

func main() {
	logger := slogger.NewLogger()

	cfg := config.MustLoad()

	dbConnection, err := postgres.NewConnection(cfg.DBString)
	if err != nil {
		panic(fmt.Errorf("failed to connect to the database: %w", err))
	}

	defer func(dbConn *sql.DB) {
		if err = postgres.CloseConnection(dbConn); err != nil {
			logger.Error("failed to close the database connection", slogger.Err(err))
			return
		}
	}(dbConnection)

	users := make([]model.User, userCount)
	for i := 0; i < userCount; i++ {
		dateOfBirth, err := time.Parse(time.DateOnly, faker.Date())
		if err != nil {
			panic(fmt.Errorf("failed to parse date of birth: %w", err))
		}

		genders := []model.Gender{model.Male, model.Female, model.Other}
		randomGenderIdx := rand.Intn(len(genders))

		cities := []string{"New York", "Los Angeles", "Chicago", "Houston", "Phoenix", "Philadelphia", "San Antonio", "San Diego", "Dallas", "San Jose"}
		randomCityIdx := rand.Intn(len(cities))

		users[i] = model.User{
			Email:       faker.Email(),
			Password:    faker.Password(),
			FirstName:   faker.FirstName(),
			LastName:    faker.LastName(),
			DateOfBirth: dateOfBirth,
			Gender:      genders[randomGenderIdx],
			Interests:   faker.Word(),
			City:        cities[randomCityIdx],
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
	}

	err = insertUsers(dbConnection, users)
	if err != nil {
		panic(fmt.Errorf("failed to insert users: %w", err))
	}

	fmt.Println("Users have been successfully inserted")
}

func insertUsers(db *sql.DB, users []model.User) error {
	for i := 0; i < len(users); i += batchSize {

		end := i + batchSize
		if end > len(users) {
			end = len(users)
		}

		query := "INSERT INTO users (email, password, first_name, last_name, date_of_birth, gender, interests, city) VALUES "

		var values []interface{}

		// Формируем строки значений ($1, $2, $3, $4, $5, $6, $7, $8), ($9, $10, $11, $12, $13, $14, $15, $16), ...
		var valueStrings []string
		for j, user := range users[i:end] {
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
				j*8+1, j*8+2, j*8+3, j*8+4, j*8+5, j*8+6, j*8+7, j*8+8))
			values = append(values, user.Email, user.Password, user.FirstName, user.LastName, user.DateOfBirth, user.Gender, user.Interests, user.City)
		}

		query += strings.Join(valueStrings, ", ")

		_, err := db.Exec(query, values...)
		if err != nil {
			panic(fmt.Errorf("failed to execute batch insert: %w", err))
		}
	}

	return nil
}
