package services

import (
	"encoding/json"
	"fmt"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/dtos"
	"io/ioutil"
	"net/http"
)

// MovieService handles movie-related operations
type MovieService struct {
	APIKey   string
	Endpoint string
}

// NewMovieService creates a new MovieService instance
func NewMovieService(apiKey, endpoint string) *MovieService {
	return &MovieService{
		APIKey:   apiKey,
		Endpoint: endpoint,
	}
}

// GetMoviesForRent retrieves movies based on a query
func (s *MovieService) GetMoviesForRent(query string) ([]dtos.MovieDTO, error) {
	url := fmt.Sprintf("%s?apikey=%s&s=%s", s.Endpoint, s.APIKey, query)
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %v", err)
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return nil, fmt.Errorf("Error decoding JSON response: %v", err)
	}

	// Extract the "Search" array from the JSON response
	searchArray, ok := jsonResponse["Search"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("Expected 'Search' array in JSON response")
	}

	// Unmarshal each movie in the "Search" array
	var result []dtos.MovieDTO
	for _, movieData := range searchArray {
		movieBytes, err := json.Marshal(movieData)
		if err != nil {
			return nil, fmt.Errorf("Error encoding movie data: %v", err)
		}

		var movie dtos.MovieDTO
		if err := json.Unmarshal(movieBytes, &movie); err != nil {
			return nil, fmt.Errorf("Error decoding movie data: %v", err)
		}

		result = append(result, movie)
	}

	return result, nil
}
