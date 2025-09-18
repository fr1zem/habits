package Postgres

import (
	"database/sql"
)

func NewPostgres(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	return db, err
}
