package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Article struct {
	gorm.Model
	ArticleID   int    `json:"article_id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	IsPublished byte   `json:"is_published"`
}

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
	DB.AutoMigrate(Article{})
	if err != nil {
		log.Fatal("Failed connection to database...", err)
	}
	fmt.Println("Connection success...")
}
