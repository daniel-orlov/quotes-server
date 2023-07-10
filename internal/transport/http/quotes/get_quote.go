package quotes

import (
	"github.com/gin-gonic/gin"
)

// GetQuote handles the request for getting a quote.
func (h *Handler) GetQuote(c *gin.Context) {
	// not implemented
	c.AbortWithStatusJSON(501, nil)
}
