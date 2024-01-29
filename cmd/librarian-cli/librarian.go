package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "librarian",
	Short: "A CLI tool for interacting with book management service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Specify a subcommand. Use 'librarian --help' for more information.")
	},
}

var getBooksCmd = &cobra.Command{
	Use:   "get-books",
	Short: "Get all books",
	Run: func(cmd *cobra.Command, args []string) {
		makeRequest("GET", "http://localhost:8080/book/", nil, nil)
	},
}

var createBookCmd = &cobra.Command{
	Use:   "create-book [title] [author] [genre] [publicationDate]",
	Short: "Create a new book",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		title, author, genre := args[0], args[1], args[2]
		url := "http://localhost:8080/book/"
		payload := map[string]string{"title": title, "author": author, "genre": genre}
		makeRequest("POST", url, nil, payload)
	},
}

var filterBooksByAuthorCmd = &cobra.Command{
	Use:   "filter-by-author",
	Short: "Filter books by author",
	Run: func(cmd *cobra.Command, args []string) {
		author, _ := cmd.Flags().GetString("author")
		if author == "" {
			fmt.Println("Author name must be provided")
			return
		}
		url := "http://localhost:8080/book/filter-by-author?author=" + url.QueryEscape(author)
		makeRequest("GET", url, nil, nil)
	},
}

var filterBooksByGenreCmd = &cobra.Command{
	Use:   "filter-by-genre",
	Short: "Filter books by genre",
	Run: func(cmd *cobra.Command, args []string) {
		genre, _ := cmd.Flags().GetString("genre")
		if genre == "" {
			fmt.Println("Genre must be provided")
			return
		}
		url := "http://localhost:8080/book/filter-by-genre?genre=" + url.QueryEscape(genre)
		makeRequest("GET", url, nil, nil)
	},
}

var getAllCollectionsCmd = &cobra.Command{
	Use:   "get-collections",
	Short: "Retrieve all collections",
	Run: func(cmd *cobra.Command, args []string) {
		url := "http://localhost:8080/collection/"
		makeRequest("GET", url, nil, nil)
	},
}

var createCollectionCmd = &cobra.Command{
	Use:   "create-collection [name]",
	Short: "Create a new collection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		url := "http://localhost:8080/collection/"
		payload := map[string]string{"name": name}
		makeRequest("POST", url, nil, payload)
	},
}

var filterCollectionCmd = &cobra.Command{
	Use:   "filter-collection [collectionId]",
	Short: "Filter books in a collection by author, genre, or publication date range",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		collectionID := args[0]
		author, _ := cmd.Flags().GetString("author")
		genre, _ := cmd.Flags().GetString("genre")
		url := fmt.Sprintf("http://localhost:8080/collection/%s/books/filter", collectionID)
		params := map[string]string{"author": author, "genre": genre}
		makeRequest("GET", url, params, nil)
	},
}

var addBookToCollectionCmd = &cobra.Command{
	Use:   "add-book [collectionId] [bookId]",
	Short: "Add a single book to a collection",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("CLI: adding book to collection...")
		collectionID := args[0]
		bookID := args[1]
		url := fmt.Sprintf("http://localhost:8080/collection/%s/add-book", collectionID)
		payload := map[string]string{"ID": bookID}
		makeRequest("POST", url, nil, payload)
	},
}

var getBooksFromCollectionCmd = &cobra.Command{
	Use:   "get-collection-books [collectionId]",
	Short: "Retrieve all books from a specific collection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		collectionID := args[0]
		url := fmt.Sprintf("http://localhost:8080/collection/%s/books", collectionID)
		makeRequest("GET", url, nil, nil)
	},
}

var filterBooksInCollectionByGenreCmd = &cobra.Command{
	Use:   "filter-collection-by-genre [collectionId] --genre=\"genreName\"",
	Short: "Filter books in a collection by genre",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		collectionID := args[0]
		genre, _ := cmd.Flags().GetString("genre")
		if genre == "" {
			fmt.Println("Genre must be provided")
			return
		}
		url := fmt.Sprintf("http://localhost:8080/collection/%s/books/filter-by-genre?genre=%s", collectionID, url.QueryEscape(genre))
		makeRequest("GET", url, nil, nil)
	},
}

var filterBooksInCollectionByAuthorCmd = &cobra.Command{
	Use:   "filter-collection-by-author [collectionId] --author=\"authorName\"",
	Short: "Filter books in a collection by author",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		collectionID := args[0]
		author, _ := cmd.Flags().GetString("author")
		if author == "" {
			fmt.Println("Author must be provided")
			return
		}
		url := fmt.Sprintf("http://localhost:8080/collection/%s/books/filter-by-author?author=%s", collectionID, url.QueryEscape(author))
		makeRequest("GET", url, nil, nil)
	},
}

func init() {
	rootCmd.AddCommand(
		getBooksCmd,
		createBookCmd,
		getAllCollectionsCmd,
		createCollectionCmd,
		addBookToCollectionCmd,
		filterCollectionCmd,
		getBooksFromCollectionCmd,
		getBooksFromCollectionCmd,
		filterBooksByAuthorCmd,
		filterBooksByGenreCmd,
		filterBooksInCollectionByGenreCmd,
		filterBooksInCollectionByAuthorCmd,
	)

	filterBooksInCollectionByAuthorCmd.Flags().String("author", "", "Author to filter the books by")
	filterBooksInCollectionByGenreCmd.Flags().String("genre", "", "Genre to filter the books by")
	filterBooksByGenreCmd.Flags().String("genre", "", "Genre to filter the books by")
	filterBooksByAuthorCmd.Flags().String("author", "", "Author name to filter the books by")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func makeRequest(method, urlStr string, params, payload map[string]string) {
	var body *bytes.Buffer
	if method == "POST" || method == "PUT" {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			fmt.Printf("Error marshaling payload to JSON: %v\n", err)
			return
		}
		body = bytes.NewBuffer(jsonData)
		fmt.Println("Sending JSON payload:", body.String())
	} else {
		body = bytes.NewBuffer(nil)
	}

	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	if method == "POST" || method == "PUT" {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}
	fmt.Printf("Response Status: %s\nBody: %s\n", resp.Status, string(responseBytes))
}
