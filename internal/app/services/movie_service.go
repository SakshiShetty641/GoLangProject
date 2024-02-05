package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/dto"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/repository"
	"io/ioutil"
	"log"
	"net/http"
)

type MovieService struct {
	APIKey          string
	Endpoint        string
	MovieRepository *repository.MovieRepository
}

func NewMovieService(apiKey, endpoint string, movieRepository *repository.MovieRepository) *MovieService {
	return &MovieService{
		APIKey:          apiKey,
		Endpoint:        endpoint,
		MovieRepository: movieRepository,
	}
}

func (s *MovieService) GetMoviesForRent(query string) (dto.Movie, error) {

	url := fmt.Sprintf("%s?apikey=%s&t=%s", s.Endpoint, s.APIKey, query)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal("HTTP request failed ")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatal("Unexpected status code: %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response body: %v", err)
	}

	var movie dto.Movie
	if err := json.Unmarshal(body, &movie); err != nil {
		return dto.Movie{}, errors.New("Error decoding movie data")
	}

	return movie, nil

}

func (s *MovieService) GetFilteredMovies(genre, actor, year string) ([]dto.Movie, error) {
	return s.MovieRepository.GetFilteredMovies(genre, actor, year)
}
