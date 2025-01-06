package db

import (
	error2 "Market/error"
	"database/sql"
	_ "github.com/lib/pq"
)

func Connect(DatabaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", DatabaseURL)
	if err != nil {
		return nil, error2.Wrap("Failed connect to DB", err)
	}

	if err := db.Ping(); err != nil {
		return nil, error2.Wrap("Error in Ping()", err)
	}

	return db, nil

}
