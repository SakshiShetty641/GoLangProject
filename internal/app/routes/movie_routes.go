package routes

import (
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/handlers"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/repository"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/services"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/db"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	apiKey := "48c1dcb7"
	endpoint := "https://www.omdbapi.com/"
	dbConn := db.CreateConnection()
	movieRepository := repository.NewMovieRepository(dbConn)
	movieService := services.NewMovieService(apiKey, endpoint)
	movieHandler := handlers.NewMovieHandler(movieService, movieRepository)
	engine.GET("/movies", movieHandler.GetMoviesForRentHandler)
}
