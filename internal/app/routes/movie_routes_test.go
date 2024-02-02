package routes

import (
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/dto"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/handlers"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/repository"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/services"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMoviesForRent(t *testing.T) {
	router := gin.Default()
	movieRepository := repository.NewOMDBMovieRepository("48c1dcb7", "http://www.omdbapi.com/")
	movieService := services.NewMovieService(movieRepository)
	movieHandler := handlers.NewMovieHandler(movieService)
	SetMovieRoutes(router, movieHandler)
	t.Run("Success", func(t *testing.T) {
		server := httptest.NewServer(router)
		defer server.Close()
		response, err := http.Get(server.URL + "/movies?query=love")
		assert.NoError(t, err)
		defer response.Body.Close()
		assert.Equal(t, http.StatusOK, response.StatusCode)
		var movies []dto.Movie
		err = json.NewDecoder(response.Body).Decode(&movies)

		expectedMovie := dto.Movie{Title: "Crazy, Stupid, Love.", Year: "2011", Poster: "https://m.media-amazon.com/images/M/MV5BMTg2MjkwMTM0NF5BMl5BanBnXkFtZTcwMzc4NDg2NQ@@._V1_SX300.jpg"}
		assert.NoError(t, err)
		expectedMovie = dto.Movie{Title: "Crazy, Stupid, Love.", Year: "2011", Poster: "https://m.media-amazon.com/images/M/MV5BMTg2MjkwMTM0NF5BMl5BanBnXkFtZTcwMzc4NDg2NQ@@._V1_SX300.jpg"}
		assert.Equal(t, expectedMovie, movies[0])
	})
	t.Run("Error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("Internal Server Error"))
			if err != nil {
				return
			}
		}))
		defer server.Close()
		movieRepository := repository.NewOMDBMovieRepository("48c1dcb7", "http://www.omdbapi.com/")
		movieService := services.NewMovieService(movieRepository)
		router := gin.Default()
		movieHandler := handlers.NewMovieHandler(movieService)
		SetMovieRoutes(router, movieHandler)
		response, err := http.Get(server.URL + "/movies?query=error_scenario")
		assert.NoError(t, err)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(response.Body)
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
	})
}
