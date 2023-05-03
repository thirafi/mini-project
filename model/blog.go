package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
