package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

func NewConnection(dbString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(95)                  // Максимальное количество открытых соединений (активных)
	db.SetMaxIdleConns(95)                  // Максимальное количество бездействующих соединений (в пуле ожидания)
	db.SetConnMaxLifetime(30 * time.Minute) // Максимальное время жизни соединения

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CloseConnection(db *sql.DB) error {
	return db.Close()
}
