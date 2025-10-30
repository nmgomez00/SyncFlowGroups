package db

import (
	// Bloque 1: Librerías estándar
	"log"
	"os"

	// Bloque 2: Librerías de terceros (sqlx, godotenv)
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	// Bloque 3: Importaciones de drivers (Importación Silenciosa)
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
	connStr := os.Getenv("DB_URL")
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln("Fallo al conectarse a la base de datos:", err)
	}
	return db

}
