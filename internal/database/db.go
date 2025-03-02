package database

import (
	"fmt"
	"log"
	"os"

	"github.com/project-Grok3_Golang/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgUser := os.Getenv("PG_USER")
	pgPassword := os.Getenv("PG_PASSWORD")
	pgDB := os.Getenv("PG_DB")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", pgHost, pgPort, pgUser, pgPassword, pgDB)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	DB.AutoMigrate(&models.User{})
}
