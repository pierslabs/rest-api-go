package dbPostgres

import (
	"fmt"
	"log"

	"github.com/pierslabs/rest-api-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func PostgresConnection() {
	cfg, err := config.LoadPostgresConfig()

	if err != nil {
		log.Fatal("Error loading PostgreSQL configuration.: %w", err)
	}

	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect to the database.: %w ❌", err)
	}

	println("Connected to the database. ✅")

}
