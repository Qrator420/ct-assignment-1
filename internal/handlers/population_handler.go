package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"strings"
	"time"

	"assignment-1/internal/config"
	"assignment-1/internal/structs"
	"assignment-1/internal/utils"
)

// FetchPopulationData handles requests for population data of a specific country
func FetchPopulationData(w http.ResponseWriter, r *http.Request) {
	// Extract country code from URL path
	countryCode := strings.TrimPrefix(r.URL.Path, "/countryinfo/v1/population/")
	if countryCode == "" {
		http.Error(w, "Missing country code", http.StatusBadRequest)
		return
	}

	// Parse optional year range parameter
	var startYear, endYear int
	var err error

	limitParam := r.URL.Query().Get("limit")
	if limitParam != "" {
		startYear, endYear, err = utils.ParseYearRange(limitParam)
		if err != nil {
			http.Error(w, "Invalid year range format. Expected format: startYear-endYear", http.StatusBadRequest)
			return
		}
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: time.Duration(config.APITimeout) * time.Second,
	}

	// First, get the country name from the REST Countries API
	resp, err := client.Get(config.BaseURLRestCountries + "alpha/" + countryCode)
	if err != nil {
		http.Error(w, "Failed to fetch country info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error from REST Countries API: "+resp.Status, resp.StatusCode)
		return
	}

	var countryData []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&countryData); err != nil {
		http.Error(w, "Failed to decode country info response", http.StatusInternalServerError)
		return
	}

	if len(countryData) == 0 {
		http.Error(w, "Country not found", http.StatusNotFound)
		return
	}

	// Extract the country name
	countryName := countryData[0]["name"].(map[string]interface{})["common"].(string)

	// Now fetch population data from CountriesNow API
	resp, err = client.Get(config.BaseURLCountriesNow + "countries/population")
	if err != nil {
		http.Error(w, "Failed to fetch population data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error from CountriesNow API: "+resp.Status, resp.StatusCode)
		return
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, "Failed to decode population data response", http.StatusInternalServerError)
		return
	}

	// Extract the data array from the response
	dataArray, ok := data["data"].([]interface{})
	if !ok {
		http.Error(w, "Unexpected response format from CountriesNow API", http.StatusInternalServerError)
		return
	}

	// Find the requested country
	var countryPopData map[string]interface{}
	for _, item := range dataArray {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		countryItemName, ok := itemMap["country"].(string)
		if !ok {
			continue
		}

		// Case-insensitive comparison
		if strings.EqualFold(countryItemName, countryName) {
			countryPopData = itemMap
			break
		}
	}

	if countryPopData == nil {
		http.Error(w, "Country population data not found", http.StatusNotFound)
		return
	}

	// Extract population data
	populationCounts, ok := countryPopData["populationCounts"].([]interface{})
	if !ok {
		http.Error(w, "Invalid population data format", http.StatusInternalServerError)
		return
	}

	// Format the response according to requirements
	var populationValues []structs.PopulationYear
	var sum int64
	count := 0

	for _, popItem := range populationCounts {
		popMap, ok := popItem.(map[string]interface{})
		if !ok {
			continue
		}

		year, ok := popMap["year"].(float64)
		if !ok {
			continue
		}

		value, ok := popMap["value"].(float64)
		if !ok {
			continue
		}

		yearInt := int(year)

		// Apply year range filter if specified
		if limitParam != "" {
			if yearInt < startYear || yearInt > endYear {
				continue
			}
		}

		// Add to population values
		populationValues = append(populationValues, structs.PopulationYear{
			Year:  yearInt,
			Value: int(value),
		})

		sum += int64(value)
		count++
	}

	// Calculate mean population
	var mean int
	if count > 0 {
		mean = int(math.Round(float64(sum) / float64(count)))
	}

	// Create response
	response := structs.PopulationResponse{
		Mean:   mean,
		Values: populationValues,
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
