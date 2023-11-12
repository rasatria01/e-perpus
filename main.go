package main

import (
	"net/http"
	"rasatria01/e-perpus/routes"
)

func main() {
	router := routes.SetupRoutes()

	http.ListenAndServe(":8080", router)
}
