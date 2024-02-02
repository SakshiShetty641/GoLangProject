package main

import (
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/handlers"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/repository"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/routes"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/services"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	apiKey := "48c1dcb7"
	endpoint := "https://www.omdbapi.com/"
	movieRepo := repository.NewOMDBMovieRepository(apiKey, endpoint)
	movieService := services.NewMovieService(movieRepo)
	movieHandler := handlers.NewMovieHandler(movieService)

	routes.SetMovieRoutes(router, movieHandler)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
