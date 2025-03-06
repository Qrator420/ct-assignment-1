package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"assignment-1/internal/config"
	"assignment-1/internal/structs"
)

// FetchCountryInfo handles requests for information about a specific country by code
func FetchCountryInfo(w http.ResponseWriter, r *http.Request) {
	countryCode := strings.TrimPrefix(r.URL.Path, "/countryinfo/v1/info/")
	if countryCode == "" {
		http.Error(w, "Missing country code", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(config.BaseURLRestCountries + "alpha/" + countryCode)
	if err != nil {
		http.Error(w, "Failed to fetch country info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var countryData []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&countryData); err != nil {
		http.Error(w, "Failed to decode response", http.StatusInternalServerError)
		return
	}

	if len(countryData) == 0 {
		http.Error(w, "Country not found", http.StatusNotFound)
		return
	}

	info := structs.CountryInfo{
		Name:       countryData[0]["name"].(map[string]interface{})["common"].(string),
		Capital:    countryData[0]["capital"].([]interface{})[0].(string),
		Region:     countryData[0]["region"].(string),
		Subregion:  countryData[0]["subregion"].(string),
		Population: int(countryData[0]["population"].(float64)),
	}

	json.NewEncoder(w).Encode(info)
}

// StatusHandler provides a simple health check endpoint
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service is running"))
}
