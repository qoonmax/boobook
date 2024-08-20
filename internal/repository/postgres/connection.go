package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func NewConnection(dbString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CloseConnection(db *sql.DB) error {
	return db.Close()
}
