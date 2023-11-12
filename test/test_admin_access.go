package test

import (
	"bytes"
	"rasatria01/e-perpus/controllers"
	"rasatria01/e-perpus/models"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBookWithAdminAccess(t *testing.T) {
	// Assuming you have an authenticated admin user
	adminUser := models.User{ID: "1", Username: "admin", Password: "adminpass", Role: "admin"}

	// Prepare a new book to create
	newBook := models.Book{Title: "Test Book"}

	// Convert newBook to JSON
	newBookJSON, _ := json.Marshal(newBook)

	req, err := http.NewRequest("POST", "/librarian/books", bytes.NewBuffer(newBookJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Add authentication header for admin access
	req.SetBasicAuth(adminUser.Username, adminUser.Password)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreateBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("CreateBook endpoint returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Add additional assertions if needed, e.g., check the response body for the created book details
}

func TestUpdateBookWithAdminAccess(t *testing.T) {
	// Assuming you have an authenticated admin user
	adminUser := models.User{ID: "1", Username: "admin", Password: "adminpass", Role: "admin"}

	// Assuming you have a book ID to update
	bookID := "123"

	// Prepare updated book details
	updatedBook := models.Book{ID: bookID, Title: "Updated Book"}

	// Convert updatedBook to JSON
	updatedBookJSON, _ := json.Marshal(updatedBook)

	req, err := http.NewRequest("PUT", "/librarian/books/"+bookID, bytes.NewBuffer(updatedBookJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Add authentication header for admin access
	req.SetBasicAuth(adminUser.Username, adminUser.Password)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.UpdateBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("UpdateBook endpoint returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Add additional assertions if needed, e.g., check the response body for the updated book details
}

func TestDeleteBookWithAdminAccess(t *testing.T) {
	// Assuming you have an authenticated admin user
	adminUser := models.User{ID: "1", Username: "admin", Password: "adminpass", Role: "admin"}

	// Assuming you have a book ID to delete
	bookID := "123"

	req, err := http.NewRequest("DELETE", "/librarian/books/"+bookID, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Add authentication header for admin access
	req.SetBasicAuth(adminUser.Username, adminUser.Password)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.DeleteBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("DeleteBook endpoint returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Add additional assertions if needed, e.g., check the response body for a success message
}
