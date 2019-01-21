package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jirawatfreedom/rest-api/handlers"
	"github.com/jirawatfreedom/rest-api/version"
)

// How to try it: go run main.go
func main() {
	log.Printf(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)
	log.Print("Starting the service...")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	r := handlers.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
