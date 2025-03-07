package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"assignment-1/internal/handlers"
)

func main() {
	// Set up routes
	http.HandleFunc("/countryinfo/v1/info/", handlers.FetchCountryInfo)
	http.HandleFunc("/countryinfo/v1/population/", handlers.FetchPopulationData)
	http.HandleFunc("/countryinfo/v1/status", handlers.StatusHandler)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Start server
	fmt.Printf("Server running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
