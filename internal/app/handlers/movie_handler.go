package handlers

import (
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MovieHandler struct {
	MovieService *services.MovieService
}

func NewMovieHandler(movieService *services.MovieService) *MovieHandler {
	return &MovieHandler{
		MovieService: movieService,
	}
}

func (h *MovieHandler) GetMoviesForRentHandler(c *gin.Context) {
	query := c.Query("query")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameter"})
		return
	}

	movies, err := h.MovieService.GetMoviesForRent(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}
