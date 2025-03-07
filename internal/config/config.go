package config

import (
	"os"
	"strconv"
	"time"
)

// API constants
var (
	BaseURLCountriesNow  = getEnvOrDefault("COUNTRIES_NOW_API", "http://129.241.150.113:3500/api/v0.1/")
	BaseURLRestCountries = getEnvOrDefault("REST_COUNTRIES_API", "http://129.241.150.113:8080/v3.1/")
	APITimeout           = getEnvAsIntOrDefault("API_TIMEOUT", 5)
)

// ServiceStartTime stores the time when the service was started
var ServiceStartTime = time.Now()

// Version of the API
const Version = "v1"

// getEnvOrDefault retrieves an environment variable or returns a default value if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsIntOrDefault retrieves an environment variable as an integer or returns a default value
func getEnvAsIntOrDefault(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
