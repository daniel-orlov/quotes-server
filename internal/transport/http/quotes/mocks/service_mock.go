package mocks

import (
	"context"

	"github.com/oklog/ulid/v2"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
)

// MockQuoteService is a mock implementation of the quotes service.
type MockQuoteService struct {
	quotes       map[ulid.ULID]model.Quote
	serviceError error
}

// NewMockQuoteService creates a new mock quotes service.
func NewMockQuoteService(quotes map[ulid.ULID]model.Quote, serviceError error) *MockQuoteService {
	return &MockQuoteService{quotes: quotes, serviceError: serviceError}
}

// GetRandomQuote returns a random quote.
func (m *MockQuoteService) GetRandomQuote(_ context.Context) (*model.Quote, error) {
	// If the service error is not nil, return it.
	if m.serviceError != nil {
		return nil, m.serviceError
	}

	// If the map of quotes is nil, create a new one.
	if m.quotes == nil {
		m.quotes = make(map[ulid.ULID]model.Quote)
	}

	// creating a slice of quotes to pick a random one
	quoteList := make([]model.Quote, 0, len(m.quotes))

	// iterating over the map of quotes and appending them to the slice
	for _, quote := range m.quotes {
		quoteList = append(quoteList, quote)
	}

	// Picking a random quote from the slice
	// Checking if the slice is empty
	if len(quoteList) == 0 {
		return nil, nil
	}

	// Since the order of map iteration is random, this is a random quote
	return &quoteList[0], nil
}
