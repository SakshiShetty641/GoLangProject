package handlers

import (
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/repository"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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
	genreParam := c.Query("genre")
	genres := strings.Split(genreParam, ",")

	// Trim spaces from each genre
	for i := range genres {
		genres[i] = strings.TrimSpace(genres[i])
	}

	actorParam := c.Query("actor")
	actors := strings.Split(actorParam, ",")

	// Trim spaces from each actor
	for i := range actors {
		actors[i] = strings.TrimSpace(actors[i])
	}

	yearParam := c.Query("year")

	// Split the year parameter into single years and ranges
	yearParts := strings.Split(yearParam, ",")
	var years []string

	for _, part := range yearParts {
		// Check if it's a range
		rangeParts := strings.Split(part, "-")
		if len(rangeParts) == 2 {
			// It's a range, add all years in the range
			startYear := rangeParts[0]
			endYear := rangeParts[1]
			for i := startYear; i <= endYear; i = strconv.Itoa(+1) {
				years = append(years, i)
			}
		} else {
			// It's a single year
			years = append(years, part)
		}
	}

	movies, err := h.MovieService.GetFilteredMovies(genres, actors, years)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
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
