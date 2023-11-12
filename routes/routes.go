package routes

import (
	"rasatria01/e-perpus/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.Login).Methods("POST")

	adminAuthMiddleware := controllers.Authenticate(controllers.AdminAccess, "admin")
	router.Handle("/admin/books", adminAuthMiddleware).Methods("GET")

	// Routes with librarian access
	librarianAuthMiddleware := controllers.Authenticate(controllers.LibrarianAccess, "librarian")
	router.Handle("/librarian/books", librarianAuthMiddleware).Methods("GET", "POST")

	// Routes with member access
	memberAuthMiddleware := controllers.Authenticate(controllers.MemberAccess, "member")
	router.Handle("/member/books", memberAuthMiddleware).Methods("GET")

	return router

}
