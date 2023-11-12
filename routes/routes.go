package routes

import (
	"rasatria01/e-perpus/controllers"
	"rasatria01/e-perpus/utils"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.Use(utils.LoggerMiddleware)
	// Books routes
	router.HandleFunc("/v1/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/v1/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/v1/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/v1/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/v1/books/{id}", controllers.DeleteBook).Methods("DELETE")

	return router
}
