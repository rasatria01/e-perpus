package controllers

import (
	"encoding/json"
	"net/http"
	"rasatria01/e-perpus/models"
	"rasatria01/e-perpus/utils"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	books     = make(map[string]models.Book)
	booksLock sync.RWMutex
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	booksLock.RLock()
	defer booksLock.RUnlock()

	bookList := make([]models.Book, 0, len(books))
	for _, book := range books {
		bookList = append(bookList, book)
	}

	utils.JsonResponse(w, http.StatusOK, bookList)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["id"]

	booksLock.RLock()
	defer booksLock.RUnlock()

	book, found := books[bookID]

	if !found {
		utils.JsonError(w, http.StatusNotFound, "Book not found")
		return
	}

	utils.JsonResponse(w, http.StatusOK, book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book

	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		utils.JsonError(w, http.StatusBadRequest, "Invalid request Payload")
		return
	}

	newBook.ID = uuid.New().String()

	booksLock.Lock()
	defer booksLock.Unlock()

	books[newBook.ID] = newBook

	utils.JsonResponse(w, http.StatusCreated, newBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["id"]

	booksLock.Lock()
	defer booksLock.Unlock()

	_, found := books[bookID]
	if !found {
		utils.JsonError(w, http.StatusNotFound, "Book not found")
		return
	}

	var UpdateBook models.Book

	err := json.NewDecoder(r.Body).Decode((&UpdateBook))

	if err != nil {
		utils.JsonError(w, http.StatusBadRequest, "Invalid Request Payload")
		return
	}

	UpdateBook.ID = bookID
	books[bookID] = UpdateBook
	utils.JsonResponse(w, http.StatusOK, UpdateBook)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["id"]

	booksLock.Lock()
	defer booksLock.Unlock()

	_, found := books[bookID]
	if !found {
		utils.JsonError(w, http.StatusNotFound, "Book not found")
		return
	}

	delete(books, bookID)

	utils.JsonResponse(w, http.StatusOK, map[string]string{"message": "Book deleted successfully"})
}
