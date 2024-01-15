package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID            int       `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	PublishedDate time.Time `json:"published_date"`
}
