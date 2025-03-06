package structs

type CountryInfo struct {
	Name       string   `json:"name"`
	Capital    string   `json:"capital"`
	Region     string   `json:"region"`
	Subregion  string   `json:"subregion"`
	Population int      `json:"population"`
	Languages  []string `json:"languages"`
}
