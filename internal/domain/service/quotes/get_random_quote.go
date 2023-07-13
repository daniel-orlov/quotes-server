package quotes

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
)

// GetRandomQuote returns a random quote.
// It checks if the cache is empty, and if so, it loads all quotes from storage and caches them.
// Then it returns a random quote from the cache.
func (s *Service) GetRandomQuote(ctx context.Context) (*model.Quote, error) {
	// Logging the call to the service.
	s.logger.Debug("getting random quote")

	// Checking if the cache is empty.
	if len(s.cache) == 0 {
		// If the cache is empty, we load all quotes from storage.
		quoteList, err := s.storage.GetQuoteList(ctx)
		// If there was an error, we return it, wrapping it with a message.
		if err != nil {
			return nil, fmt.Errorf("getting quote list: %w", err)
		}

		// Checking if the quote list is empty.
		if len(quoteList) == 0 {
			// If the quote list is empty, we return an error.
			return nil, errors.New("quote list is empty")
		}

		// If there was no error, we cache the quotes.
		s.cache = quoteList
	}

	// Pick a random quote from the cache.
	// We use the math/rand package to generate a random index.
	randIndex := rand.Intn(len(s.cache)) //nolint:gosec // We don't need a cryptographically secure random number here.

	// We use the random index to get a random quote from the cache.
	// We return the quote and nil as the error.
	quote := s.cache[randIndex]

	// Logging the result.
	s.logger.Debug("got random quote", zap.String("quote", quote.Text))

	// Returning the quote and nil as the error.
	return &quote, nil
}
