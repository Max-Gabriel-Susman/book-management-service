package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Max-Gabriel-Susman/book-management-service/internal/models"
	"github.com/Max-Gabriel-Susman/book-management-service/internal/store"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := store.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	var bookInput models.BookInput
	if err := json.NewDecoder(r.Body).Decode(&bookInput); err != nil {
		fmt.Printf("Error parsing request body: %v\n", err)
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received book input: %+v\n", bookInput)

	book := models.Book{
		Title:  bookInput.Title,
		Author: bookInput.Author,
		Genre:  bookInput.Genre,
	}

	createdBook, err := store.CreateBook(&book)
	if err != nil {
		fmt.Printf("Failed to create book: %v\n", err)
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(createdBook)
	if err != nil {
		fmt.Printf("Failed to marshal response: %v\n", err)
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func FilterBooksByAuthor(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	if author == "" {
		http.Error(w, "Author query parameter is required", http.StatusBadRequest)
		return
	}

	filteredBooks, err := store.FilterBooksByAuthor(author)
	if err != nil {
		fmt.Printf("Error filtering books by author: %v\n", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	if len(filteredBooks) == 0 {
		http.Error(w, "No books found matching the author", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(filteredBooks)
	if err != nil {
		fmt.Printf("Failed to marshal filtered books: %v\n", err)
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func FilterBooksByGenre(w http.ResponseWriter, r *http.Request) {
	genre := r.URL.Query().Get("genre")
	if genre == "" {
		http.Error(w, "Genre query parameter is required", http.StatusBadRequest)
		return
	}

	filteredBooks, err := store.FilterBooksByGenre(genre)
	if err != nil {
		fmt.Printf("Error filtering books by genre: %v\n", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	if len(filteredBooks) == 0 {
		http.Error(w, "No books found matching the genre", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(filteredBooks)
	if err != nil {
		fmt.Printf("Failed to marshal filtered books: %v\n", err)
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
