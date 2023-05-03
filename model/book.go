package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string `json:"title"`
	Creator   string `json:"creator"`
	Publisher string `json:"publisher"`
}
