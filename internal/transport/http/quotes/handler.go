// Package quotes contains the http transport for the quotes service.
package quotes

import (
	"context"

	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
)

// Service is the port for the quotes use cases.
type Service interface {
	GetRandomQuote(ctx context.Context) (*model.Quote, error)
}

// Handler is the HTTP handler for the /quotes resource.
type Handler struct {
	logger  *zap.Logger
	service Service
}

// ResourceEndpoint is the endpoint for the /quotes resource.
const ResourceEndpoint = "/quotes"

func NewHandler(logger *zap.Logger, service Service) *Handler {
	// Logging the call
	logger.Debug("creating a new quotes handler")

	return &Handler{
		logger:  logger,
		service: service,
	}
}
