package store

import (
	"github.com/Max-Gabriel-Susman/book-management-service/internal/models"
)

func CreateBook(book *models.Book) (*models.Book, error) {
	if err := db.Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func GetAllBooks() []models.Book {
	var books []models.Book
	db.Find(&books)
	return books
}

func FilterBooksByAuthor(author string) ([]models.Book, error) {
	var books []models.Book
	err := db.Where("author = ?", author).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func FilterBooksByGenre(genre string) ([]models.Book, error) {
	var books []models.Book
	err := db.Where("genre = ?", genre).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
