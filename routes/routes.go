package routes

import (
	"rasatria01/e-perpus/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

	return router

}
