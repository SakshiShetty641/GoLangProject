package services

import (
	"fmt"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/dto"
	"github.com/goccy/go-json"
	"io/ioutil"
	"net/http"
)

type MovieService struct {
	APIKey   string
	Endpoint string
}

func NewMovieService(apiKey, endpoint string) *MovieService {
	return &MovieService{
		APIKey:   apiKey,
		Endpoint: endpoint,
	}
}

func (s *MovieService) GetMoviesForRent(query string) ([]dto.Movie, error) {
	url := fmt.Sprintf("%s?apikey=%s&s=%s", s.Endpoint, s.APIKey, query)

	//url := fmt.Sprintf("%s?t=%s&apikey=%s", s.Endpoint, query, s.APIKey)

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

	searchArray, ok := jsonResponse["Search"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("Expected 'Search' array in JSON response")
	}

	var result []dto.Movie
	for _, movieData := range searchArray {
		movieBytes, err := json.Marshal(movieData)
		if err != nil {
			return nil, fmt.Errorf("Error encoding movie data: %v", err)
		}

		var movie dto.Movie
		if err := json.Unmarshal(movieBytes, &movie); err != nil {
			return nil, fmt.Errorf("Error decoding movie data: %v", err)
		}

		result = append(result, movie)
	}

	return result, nil
}
