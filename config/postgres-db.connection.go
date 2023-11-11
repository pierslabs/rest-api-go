package dbPostgres

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	PostgresPort     string
	PostgresHost     string
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	PostgresDB = os.Getenv("POSTGRES_DB")
	PostgresUser = os.Getenv("POSTGRES_USER")
	PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	PostgresPort = os.Getenv("POSTGRES_PORT")
	PostgresHost = os.Getenv("POSTGRES_HOST")

	if PostgresDB == "" || PostgresUser == "" || PostgresPassword == "" || PostgresPort == "" || PostgresHost == "" {
		log.Fatal("Alguna variable de entorno requerida no está definida en el archivo .env")
	}
}

func PostgresConnection() {

	// Crear la cadena de conexión
	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		PostgresHost, PostgresUser, PostgresPassword, PostgresDB, PostgresPort)

	_, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal("Could not connect to the database")
	} else {
		log.Println("Connect Postgres database")
	}
}
