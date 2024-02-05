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

func (h *MovieHandler) GetAllMovies(c *gin.Context) {
	query := c.Query("query")

	movie, err := h.MovieService.GetMoviesForRent(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) GetFilteredMovies(c *gin.Context) {
	genre := c.Query("genre")
	actor := c.Query("actor")
	year := c.Query("year")

	movies, err := h.MovieService.GetFilteredMovies(genre, actor, year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, movies)
}

// Story 4
func (h *MovieHandler) GetMovieDetailsByTitle(c *gin.Context) {
	title := c.Query("title")

	movie, err := h.MovieService.GetMovieDetailsByTitle(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}
