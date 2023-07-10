package quotes

import (
	"context"

	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
)

// StorageInMemory is a quote storage in memory.
type StorageInMemory struct {
	logger *zap.Logger
	db     []model.Quote
}

// NewStorageInMemory creates a new quote storage in memory.
func NewStorageInMemory(logger *zap.Logger, db []model.Quote) *StorageInMemory {
	return &StorageInMemory{logger: logger, db: db}
}

// GetQuoteList returns a list of quotes.
func (s *StorageInMemory) GetQuoteList(ctx context.Context) ([]model.Quote, error) {
	// Logging the call
	s.logger.Debug("getting quote list")

	// Checking if the context is canceled
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	// checking if the db is empty
	if len(s.db) == 0 {
		return nil, ErrDBEmpty
	}

	// returning the db and nil as the error
	return s.db, nil
}
