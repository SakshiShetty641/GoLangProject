package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(engine *gin.Engine) {
	engine.GET("/hello-world", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World Program")
	})

}
