// tests/access_denied_test.go

package test

import (
	"net/http"
	"net/http/httptest"
	"rasatria01/e-perpus/controllers"
	"rasatria01/e-perpus/models"
	"testing"
)

func TestAccessDenied(t *testing.T) {
	// Assuming you have a user with member access
	memberUser := models.User{ID: "3", Username: "member", Password: "memberpass", Role: "member"}

	req, err := http.NewRequest("GET", "/admin/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Add authentication header for member access
	req.SetBasicAuth(memberUser.Username, memberUser.Password)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.AdminAccess)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("AccessDenied test returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	// Add additional assertions if needed
}
