package database

import (
	"fmt"
	"log"
	"os"
	"test/go-rest-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnection() {
	dsn := os.Getenv("DATABASE_CREDENTIALS")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB.AutoMigrate(
		models.Article{},
		models.Image{},
		models.Video{},
	)
	if err != nil {
		log.Fatal("Failed connection to database...", err)
	}
	fmt.Println("Connection success...")
}
