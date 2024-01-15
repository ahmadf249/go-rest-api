package database

import (
	"fmt"
	"log"
	"test/go-rest-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbName := "postgres"
	dbUser := "postgres"
	password := "root"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(
		models.Article{},
		models.Image{},
	)
	if err != nil {
		log.Fatal("Failed connection to database...", err)
	}
	fmt.Println("Connection success...")
}
