package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {
	connStr := "user=postgres dbname=kokoro password=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}
