package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"assignment-1/internal/config"
)

// FetchPopulationData handles requests for population data of a specific country
func FetchPopulationData(w http.ResponseWriter, r *http.Request) {
	country := strings.TrimPrefix(r.URL.Path, "/countryinfo/v1/population/")
	if country == "" {
		http.Error(w, "Missing country name", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(config.BaseURLCountriesNow + "countries/population")
	if err != nil {
		http.Error(w, "Failed to fetch population data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, "Failed to decode response", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}
