package structs

// CountryInfo represents general country information
type CountryInfo struct {
	Name       string   `json:"name"`
	Capital    string   `json:"capital"`
	Region     string   `json:"region"`
	Subregion  string   `json:"subregion"`
	Population int      `json:"population"`
	Languages  []string `json:"languages"`
}

// PopulationResponse represents the structure required for population endpoint responses
type PopulationResponse struct {
	Mean   int              `json:"mean"`
	Values []PopulationYear `json:"values"`
}

// PopulationYear represents population data for a specific year
type PopulationYear struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

// StatusResponse represents the structure for status endpoint responses
type StatusResponse struct {
	CountriesNowAPI  interface{} `json:"countriesnowapi"`
	RestCountriesAPI interface{} `json:"restcountriesapi"`
	Version          string      `json:"version"`
	Uptime           int64       `json:"uptime"`
}
