package routes

import (
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoute(engine *gin.Engine) {
	engine.GET("/hello-world", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World Program")
	})

}

func SetMovieRoutes(router *gin.Engine, movieHandler *handlers.MovieHandler) {
	//db.CreateConnection()
	router.GET("/movies", movieHandler.GetMoviesForRentHandler)
}
