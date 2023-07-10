package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for GET /quotes/random
	router.GET("/quotes/random", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	// Start the server and listen on port 8080
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
