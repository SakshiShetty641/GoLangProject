package main

import (
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	routes.RegisterRoute(engine)
	engine.Run(":8080")
}
