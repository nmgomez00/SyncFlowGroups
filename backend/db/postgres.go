package db

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
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
	if connStr == "" {
		log.Fatalln("FATAL: DB_URL variable de entorno no configurada.")
	}

	// 1. Parsear la cadena de conexión de Supabase
	config, err := pgx.ParseConfig(connStr)
	if err != nil {
		log.Fatalln("Fallo al parsear la URL de la base de datos:", err)
	}

	// 2. ⚠️ WORKAROUND: Forzar la resolución a IPv4
	// Creamos un DialFunc personalizado.
	dialer := &net.Dialer{
		Timeout:   5 * time.Second, // Establecemos un timeout razonable
		KeepAlive: 5 * time.Second,
		// Usamos el Resolutor de Go para un manejo de DNS más estable
		Resolver: &net.Resolver{
			PreferGo: true,
		},
	}

	// Asignamos una función a DialFunc que ignora el protocolo de red
	// sugerido y siempre intenta usar 'tcp4' (solo IPv4)
	config.DialFunc = func(ctx context.Context, network, addr string) (net.Conn, error) {
		// network será 'tcp', pero lo forzamos a 'tcp4'
		return dialer.DialContext(ctx, "tcp4", addr)
	}

	// 3. Abrir la conexión usando el driver pgx configurado
	db := stdlib.OpenDB(*config)

	// 4. Usar sqlx.OpenDB para envolver la conexión
	dbx := sqlx.NewDb(db, "pgx")

	// 5. Probar la conexión (ping)
	if err = dbx.Ping(); err != nil {
		log.Fatalln("Fallo al hacer ping (conexión fallida). Verifica PASS y SSLmode:", err)
	}

	log.Println("¡Conexión a la base de datos establecida y forzada a IPv4!")
	return dbx
}
