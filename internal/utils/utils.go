package utils

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ParseYearRange parses a string like "2010-2015" into start and end years
func ParseYearRange(yearRange string) (int, int, error) {
	parts := strings.Split(yearRange, "-")
	if len(parts) != 2 {
		return 0, 0, http.ErrNotSupported
	}

	startYear, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	endYear, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return startYear, endYear, nil
}

// CalculateUptime returns the uptime in seconds since the service started
func CalculateUptime(startTime time.Time) int64 {
	return int64(time.Since(startTime).Seconds())
}

// CheckServiceStatus makes a request to check service availability
func CheckServiceStatus(url string) interface{} {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "Service unreachable"
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
