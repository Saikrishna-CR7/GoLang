package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const(
	DB_USER = "postgres"
	DB_PASSWORD ="postgres"
	DB_NAME = "PRODUCT" 
)

func ConnectDB() *sql.DB {
	connection := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err)
	}

	return db
}