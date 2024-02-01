package main

import (
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/handlers"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/routes"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin
	router := gin.Default()

	movieService := services.NewMovieService("48c1dcb7", "http://www.omdbapi.com/")
	movieHandler := handlers.NewMovieHandler(movieService)

	routes.SetMovieRoutes(router, movieHandler)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
