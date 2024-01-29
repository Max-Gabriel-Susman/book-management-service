package models

import "github.com/jinzhu/gorm"

type Collection struct {
	gorm.Model
	Name  string `json:"name"`
	Books []Book `gorm:"many2many:collection_books;" json:"books"`
}
