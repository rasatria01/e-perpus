package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rasatria01/e-perpus/controllers"
	"rasatria01/e-perpus/models"
	"testing"
)

func TestLogin(t *testing.T) {
	loginData := map[string]string{"username": "admin", "password": "adminpass"}
	loginJSON, _ := json.Marshal(loginData)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginJSON))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.Login)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Login endpoint returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestAdminAccess(t *testing.T) {
	// Assuming you have a function to authenticate and return an admin user
	adminUser := models.User{ID: "1", Username: "admin", Password: "adminpass", Role: "admin"}

	req, err := http.NewRequest("GET", "/admin/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Add authentication header
	req.SetBasicAuth(adminUser.Username, adminUser.Password)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.AdminAccess)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("AdminAccess endpoint returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Add additional assertions if needed
}
