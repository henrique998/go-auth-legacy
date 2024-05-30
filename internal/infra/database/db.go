package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToDb() *sql.DB {
	var (
		driver   = os.Getenv("DB_DRIVER")
		user     = os.Getenv("DB_USER")
		pass     = os.Getenv("DB_PASS")
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		database = os.Getenv("DB_NAME")
	)

	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user, pass, host, port, database,
	)

	db, err := sql.Open(driver, connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
