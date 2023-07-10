package quotes

import (
	"context"
	"errors"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
)

// GetRandomQuote returns a random quote.
func (s *Service) GetRandomQuote(_ context.Context) (*model.Quote, error) {
	return nil, errors.New("not implemented")
}
