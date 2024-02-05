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
	movieService := services.NewMovieService(apiKey, endpoint, movieRepository)
	movieHandler := handlers.NewMovieHandler(movieService, movieRepository)
	group := engine.Group("/api/v1")
	{
		group.GET("/movies", movieHandler.GetAllMovies)
		group.GET("/movies/filteredMovies", movieHandler.GetFilteredMovies)
		group.GET("/movies/movieDetails", movieHandler.GetMovieDetailsByTitle)

	}

}
