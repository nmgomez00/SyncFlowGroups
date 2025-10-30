package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Database *sqlx.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	Database = Connect()
}

func Connect() *sqlx.DB {
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("SSL_MODE"),
	)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln("Fallo al conectarse a la base de datos:", err)
	}
	return db

}
