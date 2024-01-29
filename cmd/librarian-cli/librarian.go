package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func init() {
	rootCmd.AddCommand(
		getBooksCmd,
		createBookCmd,
	)
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
