package crmapi

import (
	"database/sql"
	"log"
	"os"

	_"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


func ConnectDB() *sql.DB {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
