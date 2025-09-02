package database

import (
	"tonix/backend/logging"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var log = logging.LoggerWithOrigin("database.go")

func Connect(dsn string) (*sqlx.DB, error) {
	var db *sqlx.DB

	db, err := sqlx.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
