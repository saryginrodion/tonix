package database

import (
	"database/sql"
	"tonix/backend/logging"

	_ "github.com/jackc/pgx/stdlib"
)

var log = logging.LoggerWithOrigin("database.go")

func Connect(dsn string) (*sql.DB, error) {
	var db *sql.DB

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
