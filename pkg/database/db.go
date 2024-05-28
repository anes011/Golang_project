package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	connStr := "postgres://default:unI2Ax5lDLZv@ep-plain-fire-a2fjkizw.eu-central-1.aws.neon.tech:5432/verceldb?sslmode=require"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return db, err
	}

	return db, nil
}
