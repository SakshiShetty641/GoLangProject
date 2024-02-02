package services

import (
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/dto"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/repository"
)

type MovieService struct {
	Repository repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) *MovieService {
	return &MovieService{
		Repository: repo,
	}
}

func (s *MovieService) GetMoviesForRent(query string) ([]dto.Movie, error) {
	return s.Repository.GetMoviesForRent(query)
}
