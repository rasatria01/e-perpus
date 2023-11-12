package controllers

import (
	"encoding/json"
	"rasatria01/e-perpus/models"
	"rasatria01/e-perpus/utils"

	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var books []models.Book

// GetBooks returns the list of all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJSON(w, http.StatusOK, books)
}

// GetBook returns a single book by ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	for _, book := range books {
		if book.ID == bookID {
			utils.ResponseJSON(w, http.StatusOK, book)
			return
		}
	}
	utils.ResponseError(w, http.StatusNotFound, "Book not found")
}

// CreateBook adds a new book to the collection
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid request Payload")
		return
	}
	// Assign a unique ID (you may want to use a more robust ID generation mechanism)
	newBook.ID = uuid.New().String() // Replace with a proper ID assignment logic

	books = append(books, newBook)

	utils.ResponseJSON(w, http.StatusCreated, newBook)
}

// UpdateBook updates an existing book by ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	var updatedBook models.Book
	json.NewDecoder(r.Body).Decode(&updatedBook)

	for i, book := range books {
		if book.ID == bookID {
			updatedBook.ID = book.ID
			books[i] = updatedBook
			utils.ResponseJSON(w, http.StatusOK, updatedBook)
			return
		}
	}
	utils.ResponseError(w, http.StatusNotFound, "Book not Found")
}

// DeleteBook removes a book from the collection by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	for i, book := range books {
		if book.ID == bookID {
			// Remove the book from the slice
			books = append(books[:i], books[i+1:]...)
			utils.ResponseJSON(w, http.StatusNoContent, nil)
			return
		}
	}
	utils.ResponseError(w, http.StatusNotFound, "Book not Found")
}
