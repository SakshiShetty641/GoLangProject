package routes

import (
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/handlers"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/repository"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	apiKey := "48c1dcb7"
	endpoint := "https://www.omdbapi.com/"
	movieRepository := repository.NewOMDBMovieRepository(apiKey, endpoint)
	movieService := services.NewMovieService(movieRepository)
	movieHandler := handlers.NewMovieHandler(movieService)
	engine.GET("/movies", movieHandler.GetMoviesForRentHandler)
}
