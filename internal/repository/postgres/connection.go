package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func NewConnection(dbString string) (*sql.DB, error) {
	log.Println(dbString)
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseConnection(db *sql.DB) error {
	return db.Close()
}
