package http

import (
	"github.com/gin-gonic/gin"

	"github.com/daniel-orlov/quotes-server/internal/transport/http/quotes"
)

// NewRouter creates a new HTTP router.
func NewRouter(quotesHandler *quotes.Handler) *gin.Engine {
	// Initialize Gin router
	r := gin.Default()

	// Initialize an API version group
	v1 := r.Group("/v1")

	{
		// Initialize quotes group
		quoteGroup := v1.Group(quotes.ResourceEndpoint)
		{
			// Initialize quotes endpoints
			quoteGroup.GET("", quotesHandler.GetQuote)
		}
	}

	// Return router
	return r
}
