package store

import (
	"log"

	"github.com/Max-Gabriel-Susman/book-management-service/internal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	log.Println("Connecting to database...")
	d, err := gorm.Open("mysql", "admin:password@/book_management_service_database?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

func Init() {
	Connect()
	GetDB()
	db.AutoMigrate(&models.Book{}, &models.Collection{})
}
