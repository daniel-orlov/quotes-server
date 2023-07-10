package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/pkg/logging"
)

func main() {
	// Create a new logger
	logger := logging.Logger("json", "debug")

	// Create a new Gin router
	router := gin.Default()

	// Define a route for GET /quotes/random
	router.GET("/quotes/random", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	// Start the server and listen on port 8080
	err := router.Run(":8080")
	if err != nil {
		logger.Fatal("running server failed", zap.Error(err))
	}
}
