package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ArticleID   int    `json:"article_id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	IsPublished byte   `json:"is_published"`
}
