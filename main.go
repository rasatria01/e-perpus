package main

import (
	"fmt"
	"log"
	"net/http"
	"rasatria01/e-perpus/routes"
)

func main() {
	router := routes.SetupRouter()

	port := "8080"

	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
