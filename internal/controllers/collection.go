package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Max-Gabriel-Susman/book-management-service/internal/models"
	"github.com/Max-Gabriel-Susman/book-management-service/internal/store"
	"github.com/Max-Gabriel-Susman/book-management-service/internal/web"
	"github.com/gorilla/mux"
)

func CreateCollection(w http.ResponseWriter, r *http.Request) {
	createCollection := &models.Collection{}
	web.ParseBody(r, createCollection)
	c := store.CreateCollection(createCollection)
	res, _ := json.Marshal(c)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllCollections(w http.ResponseWriter, r *http.Request) {
	collections := store.GetAllCollections()
	res, _ := json.Marshal(collections)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddBookToCollection(w http.ResponseWriter, r *http.Request) {
	log.Println("Controller: adding book to collection...")
	vars := mux.Vars(r)
	collectionID := vars["collectionId"]

	var payload struct {
		BookID string `json:"ID"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	bookID, err := strconv.Atoi(payload.BookID)
	if err != nil {
		log.Println("failure to convert bookID command line argument from string to int")
	}

	if err := store.AddBookToCollection(collectionID, bookID); err != nil {
		http.Error(w, fmt.Sprintf("Error adding book to collection: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Book added to collection successfully")
}

func GetBooksFromCollection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collectionID, err := strconv.Atoi(vars["collectionId"])
	if err != nil {
		http.Error(w, "Invalid collection ID", http.StatusBadRequest)
		return
	}

	books, err := store.GetBooksFromCollection(int64(collectionID))
	if err != nil {
		http.Error(w, "Error retrieving books from collection", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Error marshaling books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func FilterCollectionBooks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collectionID := vars["collectionId"]
	author := r.URL.Query().Get("author")
	genre := r.URL.Query().Get("genre")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	filteredBooks := store.FilterCollectionBooks(collectionID, author, genre, startDate, endDate)
	res, _ := json.Marshal(filteredBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func FilterBooksInCollectionByGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collectionID := vars["collectionId"]
	genre := r.URL.Query().Get("genre")

	filteredBooks, err := store.FilterBooksInCollectionByGenre(collectionID, genre)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	if len(filteredBooks) == 0 {
		http.Error(w, "No books found matching the criteria", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(filteredBooks)
	if err != nil {
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func FilterBooksInCollectionByAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collectionID := vars["collectionId"]
	author := r.URL.Query().Get("author")

	filteredBooks, err := store.FilterBooksInCollectionByAuthor(collectionID, author)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	if len(filteredBooks) == 0 {
		http.Error(w, "No books found matching the criteria", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(filteredBooks)
	if err != nil {
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
