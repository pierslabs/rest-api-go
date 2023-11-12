package dbPostgres

import (
	"fmt"

	"github.com/pierslabs/rest-api-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func PostgresConnection() (*gorm.DB, error) {
	cfg, err := config.LoadPostgresConfig()
	if err != nil {
		return nil, fmt.Errorf("Error loading PostgreSQL configuration.: %w", err)
	}

	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the database.: %w ❌", err)
	}

	return db, nil
}
