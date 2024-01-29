package routes

import (
	"github.com/Max-Gabriel-Susman/book-management-service/internal/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/collection/", controllers.CreateCollection).Methods("POST")
	router.HandleFunc("/collection/{collectionId}/add-book", controllers.AddBookToCollection).Methods("POST")
	router.HandleFunc("/collection/", controllers.GetAllCollections).Methods("GET")
	router.HandleFunc("/book/filter-by-author", controllers.FilterBooksByAuthor).Methods("GET")
	router.HandleFunc("/book/filter-by-genre", controllers.FilterBooksByGenre).Methods("GET")
	router.HandleFunc("/collection/{collectionId}/books/filter", controllers.FilterCollectionBooks).Methods("GET")
	router.HandleFunc("/collection/{collectionId}/books", controllers.GetBooksFromCollection).Methods("GET")
	router.HandleFunc("/collection/{collectionId}/books/filter-by-genre", controllers.FilterBooksInCollectionByGenre).Methods("GET")
	router.HandleFunc("/collection/{collectionId}/books/filter-by-author", controllers.FilterBooksInCollectionByAuthor).Methods("GET")
}
