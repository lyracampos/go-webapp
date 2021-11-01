package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DatabaseConnect() *sql.DB {
	connection := "user=postgres dbname=postgres password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
