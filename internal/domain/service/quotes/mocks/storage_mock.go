package mocks

import (
	"context"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
)

// MockQuoteStorage is a mock for quote storage.
type MockQuoteStorage struct {
	quotes       []model.Quote
	storageError error
}

// NewMockQuoteStorage creates a new mock for quote storage.
func NewMockQuoteStorage(quotes []model.Quote, storageError error) *MockQuoteStorage {
	return &MockQuoteStorage{
		quotes:       quotes,
		storageError: storageError,
	}
}

// GetQuoteList returns a list of quotes.
func (m *MockQuoteStorage) GetQuoteList(_ context.Context) ([]model.Quote, error) {
	// Check if there was an error passed to the mock.
	if m.storageError != nil {
		// Return the error.
		return nil, m.storageError
	}

	// Return the quotes.
	return m.quotes, nil
}
