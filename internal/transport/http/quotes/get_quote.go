package quotes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
)

// GetQuote handles the request for getting a quote.
func (h *Handler) GetQuote(c *gin.Context) {
	// Call the service.
	quote, err := h.service.GetRandomQuote(c.Request.Context())
	// Handle the error.
	if err != nil {
		// Log the actual error.
		h.logger.Error("failed to get quote", zap.Error(err))

		// Return a generic error to the client.
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Quote{})

		// Exit the function.
		return
	}

	if quote == nil {
		// Return a generic error to the client.
		c.AbortWithStatusJSON(http.StatusNotFound, model.Quote{})

		// Exit the function.
		return
	}

	// Return the quote to the client.
	c.JSON(http.StatusOK, quote)
}
