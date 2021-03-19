package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Dependencies

	// Application routes
	engine := gin.New()
	//apiRouter := engine.Group("/api")

	engine.Run(":8080")
}
