package main

import (
	"fmt"
	"log"
	"net/http"

	"assignment-1/internal/handlers"
)

func main() {
	http.HandleFunc("/countryinfo/v1/info/", handlers.FetchCountryInfo)
	http.HandleFunc("/countryinfo/v1/population/", handlers.FetchPopulationData)
	http.HandleFunc("/countryinfo/v1/status", handlers.StatusHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
