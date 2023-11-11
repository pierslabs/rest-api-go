package dbPostgres

import (
	"fmt"
	"log"

	"github.com/pierslabs/rest-api-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresConnection() (*gorm.DB, error) {

	cfg, err := config.LoadPostgresConfig()
	if err != nil {
		log.Fatal("Error cargando la configuración de PostgreSQL:", err)
	}

	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
		return nil, err
	}

	return db, nil
}
