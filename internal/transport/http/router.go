package http

import (
	"github.com/gin-gonic/gin"

	"github.com/daniel-orlov/quotes-server/internal/transport/http/quotes"
)

// NewRouter creates a new HTTP router.
// It accepts a quotes handler and a list of global middlewares, which will be applied to all routes.
// TODO: refactor to accept a map of handlers to middleware lists for extensibility.
func NewRouter(quotesHandler *quotes.Handler, globalMWs ...gin.HandlerFunc) *gin.Engine {
	// Initialize Gin router
	r := gin.Default()

	// Initialize an API version group
	// Add global middlewares
	v1 := r.Group("/v1", globalMWs...)

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
