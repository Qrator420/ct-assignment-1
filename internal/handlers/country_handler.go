package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

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

	// Parse optional limit parameter
	limitStr := r.URL.Query().Get("limit")
	limit := 10 // Default limit
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit <= 0 {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
		limit = parsedLimit
	}

	// Create an HTTP client with timeout
	client := &http.Client{
		Timeout: time.Duration(config.APITimeout) * time.Second,
	}

	resp, err := client.Get(config.BaseURLRestCountries + "alpha/" + countryCode)
	if err != nil {
		http.Error(w, "Failed to fetch country info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error from external API: "+resp.Status, resp.StatusCode)
		return
	}

	var countryData []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&countryData); err != nil {
		http.Error(w, "Failed to decode response", http.StatusInternalServerError)
		return
	}

	if len(countryData) == 0 {
		http.Error(w, "Country not found", http.StatusNotFound)
		return
	}

	// Extract languages as array of strings with limit applied
	var languages []string
	if langs, ok := countryData[0]["languages"].(map[string]interface{}); ok {
		for _, lang := range langs {
			if langStr, ok := lang.(string); ok {
				languages = append(languages, langStr)
				// Apply limit on languages
				if len(languages) >= limit {
					break
				}
			}
		}
	}

	info := structs.CountryInfo{
		Name:       countryData[0]["name"].(map[string]interface{})["common"].(string),
		Capital:    countryData[0]["capital"].([]interface{})[0].(string),
		Region:     countryData[0]["region"].(string),
		Subregion:  countryData[0]["subregion"].(string),
		Population: int(countryData[0]["population"].(float64)),
		Languages:  languages,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}
