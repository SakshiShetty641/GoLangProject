package handlers

import (
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/repository"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MovieHandler struct {
	MovieService    *services.MovieService
	MovieRepository *repository.MovieRepository
}

func NewMovieHandler(movieService *services.MovieService, movieRepository *repository.MovieRepository) *MovieHandler {
	return &MovieHandler{
		MovieService:    movieService,
		MovieRepository: movieRepository,
	}
}

func (h *MovieHandler) GetMoviesForRentHandler(c *gin.Context) {
	query := c.Query("query")

	// Call the movie service to get movies based on the query
	movies, err := h.MovieService.GetMoviesForRent(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.MovieRepository.SaveMovie(movies)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

//func (h *MovieHandler) GetMoviesForRentHandler(c *gin.Context) {
//	query := c.Query("query")
//
//	if query == "" {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameter"})
//		return
//	}
//
//	movies, err := h.MovieService.GetMoviesForRent(query)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, movies)
//}
