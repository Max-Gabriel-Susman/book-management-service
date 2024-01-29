package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type BookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

type Book struct {
	gorm.Model
	Title       string       `json:"title" gorm:"column:title"`
	Genre       string       `json:"genre" gorm:"column:genre"`
	Author      string       `json:"author" gorm:"column:author"`
	Collections []Collection `gorm:"many2many:collection_books;" json:"-"`
}
