package controllers

import (
	"encoding/json"
	"net/http"
	"rasatria01/e-perpus/models"
	"rasatria01/e-perpus/utils"
	"sync"
)

var (
	users    = make(map[string]models.User)
	userLock sync.RWMutex
)

func init() {
	admin := models.User{ID: "1", Username: "admin", Password: "adminpass", Role: "admin"}
	librarian := models.User{ID: "2", Username: "librarian", Password: "librarianpass", Role: "librarian"}
	member := models.User{ID: "3", Username: "member", Password: "memberpass", Role: "member"}

	users[admin.ID] = admin
	users[librarian.ID] = librarian
	users[member.ID] = member
}

func AuthenticateUser(username, password string) (models.User, bool) {
	userLock.RLock()
	defer userLock.RUnlock()

	for _, user := range users {
		if user.Username == username && user.Password == password {
			return user, true
		}
	}
	return models.User{}, false
}

func GetUserByID(userID string) (models.User, bool) {
	userLock.RLock()
	defer userLock.RUnlock()

	user, found := users[userID]
	return user, found
}

func Authenticate(next http.HandlerFunc, requiredRole string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok {
			utils.JsonError(w, http.StatusUnauthorized, "Unauthorized: Missing or invalid credentials")
			return
		}

		user, authenticated := AuthenticateUser(username, password)
		if !authenticated || user.Role != requiredRole {
			utils.JsonError(w, http.StatusUnauthorized, "Unauthorized: Invalid credentials or insufficient permissions")
			return
		}
		next(w, r)
	}
}
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		utils.JsonError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, authenticated := AuthenticateUser(credentials.Username, credentials.Password)
	if !authenticated {
		utils.JsonError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	utils.JsonResponse(w, http.StatusOK, user)
}

func AdminAccess(w http.ResponseWriter, r *http.Request) {
	// Access granted to admin
	utils.JsonResponse(w, http.StatusOK, "Admin access granted")
}

func LibrarianAccess(w http.ResponseWriter, r *http.Request) {
	// Access granted to librarian
	utils.JsonResponse(w, http.StatusOK, "Librarian access granted")
}

func MemberAccess(w http.ResponseWriter, r *http.Request) {
	// Access granted to member
	utils.JsonResponse(w, http.StatusOK, "Member access granted")
}
