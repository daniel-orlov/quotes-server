package quotes

import (
	"context"

	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
)

// Storage is a port for quotes storage.
type Storage interface {
	GetQuoteList(ctx context.Context) ([]model.Quote, error)
}

// Service is a quote service.
type Service struct {
	logger  *zap.Logger
	storage Storage
	cache   []model.Quote
}

// NewService creates a new quote service.
func NewService(logger *zap.Logger, storage Storage) *Service {
	// Logging the call
	logger.Debug("creating a new quote service")

	return &Service{logger: logger, storage: storage}
}
