package quotes

import (
	"context"
	"errors"

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
func (s *StorageInMemory) GetQuoteList(_ context.Context) ([]model.Quote, error) {
	return nil, errors.New("not implemented")
}
