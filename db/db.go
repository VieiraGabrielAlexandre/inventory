package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectDB() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=root dbname=inventory host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
