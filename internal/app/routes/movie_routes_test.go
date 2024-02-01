package routes

import (
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/dtos"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/handlers"
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
	movieService := services.NewMovieService("48c1dcb7", "http://www.omdbapi.com/")
	movieHandler := handlers.NewMovieHandler(movieService)
	SetMovieRoutes(router, movieHandler)
	t.Run("Success", func(t *testing.T) {
		server := httptest.NewServer(router)
		defer server.Close()
		response, err := http.Get(server.URL + "/movies?query=love")
		assert.NoError(t, err)
		defer response.Body.Close()
		assert.Equal(t, http.StatusOK, response.StatusCode)
		var movies []dtos.MovieDTO
		err = json.NewDecoder(response.Body).Decode(&movies)
		assert.NoError(t, err)
		expectedMovie := dtos.MovieDTO{Title: "Crazy, Stupid, Love.", Year: "2011", Poster: "https://m.media-amazon.com/images/M/MV5BMTg2MjkwMTM0NF5BMl5BanBnXkFtZTcwMzc4NDg2NQ@@._V1_SX300.jpg"}
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
		movieService := services.NewMovieService("48c1dcb7", server.URL)
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