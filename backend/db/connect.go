package db

import (
	"fmt"
	"gin-quickstart/models"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to DB: %w", err)
	}
	DB = db
	if err := db.AutoMigrate(
		&models.Article{},
		&models.Link{},
		&models.Cache{},
	); err != nil {
		return fmt.Errorf("failed to migrate DB: %w", err)
	}

	return nil
}
