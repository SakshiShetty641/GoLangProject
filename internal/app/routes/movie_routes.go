package routes

import (
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/handlers"
	"github.com/gin-gonic/gin"
)

func SetMovieRoutes(router *gin.Engine, movieHandler *handlers.MovieHandler) {
	router.GET("/movies", movieHandler.GetMoviesForRentHandler)
}
