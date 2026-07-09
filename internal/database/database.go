package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jvbenetti/go-financial-core/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set in .env")
	}

	env := os.Getenv("APP_ENV")
	logLevel := logger.Info
	if env == "production" {
		logLevel = logger.Error // In production log only errors
	}

	// Connect the database and save the local var "db"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // Compatible with Supabase/PgBouncer
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

}
