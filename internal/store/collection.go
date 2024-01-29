package store

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Max-Gabriel-Susman/book-management-service/internal/models"
	"github.com/jinzhu/gorm"
)

func CreateCollection(collection *models.Collection) *models.Collection {
	db.NewRecord(collection)
	db.Create(&collection)
	return collection
}

func GetAllCollections() []models.Collection {
	var collections []models.Collection
	db.Preload("Books").Find(&collections)
	return collections
}

func FilterCollectionBooks(collectionID, author, genre, startDate, endDate string) []models.Book {
	var books []models.Book
	query := db.Table("books").
		Joins("JOIN collection_books ON books.id = collection_books.book_id").
		Where("collection_books.collection_id = ?", collectionID)

	if author != "" {
		query = query.Where("books.author = ?", author)
	}
	if genre != "" {
		query = query.Where("books.genre = ?", genre)
	}
	if startDate != "" && endDate != "" {
		startDateParsed, _ := time.Parse("2006-01-02", startDate)
		endDateParsed, _ := time.Parse("2006-01-02", endDate)
		query = query.Where("books.created_at BETWEEN ? AND ?", startDateParsed, endDateParsed)
	}
	query.Find(&books)
	return books
}

func GetBooksFromCollection(collectionID int64) ([]models.Book, error) {
	var books []models.Book
	result := db.Table("books").
		Joins("join collection_books on books.id = collection_books.book_id").
		Where("collection_books.collection_id = ?", collectionID).
		Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func AddBookToCollection(collectionID string, bookID int) error {
	log.Println("Store: adding book to a collection...")
	collID, err := strconv.ParseInt(collectionID, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid collection ID: %v", err)
	}

	tx := db.Begin()

	var book models.Book
	if tx.First(&book, bookID).RecordNotFound() {
		tx.Rollback()
		return fmt.Errorf("book with ID %d not found", bookID)
	}

	if err := tx.Model(&models.Collection{Model: gorm.Model{ID: uint(collID)}}).Association("Books").Append(&book).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to add book to collection: %v", err)
	}

	return tx.Commit().Error
}

func FilterBooksInCollectionByGenre(collectionID string, genre string) ([]models.Book, error) {
	var books []models.Book

	collID, err := strconv.ParseInt(collectionID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid collection ID format: %v", err)
	}

	err = db.Joins("JOIN collection_books ON books.id = collection_books.book_id").
		Where("collection_books.collection_id = ?", collID).
		Where("books.genre = ?", genre).
		Find(&books).Error

	if err != nil {
		return nil, fmt.Errorf("error querying books by genre in collection: %v", err)
	}

	return books, nil
}

func FilterBooksInCollectionByAuthor(collectionID string, author string) ([]models.Book, error) {
	var books []models.Book

	collID, err := strconv.ParseInt(collectionID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid collection ID format: %v", err)
	}

	err = db.Joins("JOIN collection_books ON books.id = collection_books.book_id").
		Where("collection_books.collection_id = ?", collID).
		Where("books.author = ?", author).
		Find(&books).Error

	if err != nil {
		return nil, fmt.Errorf("error querying books by author in collection: %v", err)
	}

	return books, nil
}
