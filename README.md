# Assignment-1

## Description
A RESTful API built in Go that provides country information, population data, and API status monitoring using external services.

### Features
- Retrieve general country info (name, capital, region, languages)
- Fetch population data with optional year filtering

## Installation

### Prerequisites
- Go 1.18+
- Internet connection for API access

### Clone the repository

git clone https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2025-workspace/torbjset/assignment-1.git


### Run the application
```sh
go run main.go
```
API runs at **http://localhost:8080**

### Fetch Country Info
```
GET /countryinfo/v1/info/{countryCode}
```

Response:
```json
{
  "name": "Norway",
  "capital": "Oslo",
  "region": "Europe",
  "population": 5379475,
  "languages": ["Norwegian", "Sami"]
}
```

### Fetch Population Data
```
GET /countryinfo/v1/population/{countryCode}
```

### API Status Check
```
GET /countryinfo/v1/status
```

## Authors
- **Torbjørn Seth**