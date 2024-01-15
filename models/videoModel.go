package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	ID          int    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
