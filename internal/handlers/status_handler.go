package handlers

import (
	"encoding/json"
	"net/http"

	"assignment-1/internal/config"
	"assignment-1/internal/structs"
	"assignment-1/internal/utils"
)

// StatusHandler provides diagnostic information about dependent services
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// Check the status of dependent services
	countriesNowStatus := utils.CheckServiceStatus(config.BaseURLCountriesNow)
	restCountriesStatus := utils.CheckServiceStatus(config.BaseURLRestCountries)

	// Calculate uptime
	uptime := utils.CalculateUptime(config.ServiceStartTime)

	// Create response
	status := structs.StatusResponse{
		CountriesNowAPI:  countriesNowStatus,
		RestCountriesAPI: restCountriesStatus,
		Version:          config.Version,
		Uptime:           uptime,
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
