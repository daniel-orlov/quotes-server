package quotes_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
	"github.com/daniel-orlov/quotes-server/internal/domain/service/quotes"
	"github.com/daniel-orlov/quotes-server/internal/domain/service/quotes/mocks"
)

func TestService_GetRandomQuote_Success(t *testing.T) {
	// Prepare test data.
	quoteList := []model.Quote{
		{
			ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
			Text:   "People say nothing is impossible, but I do nothing every day.",
			Author: "A. A. Milne",
		},
		{
			ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
			Text:   "The best vitamin to be a happy person is B1.",
			Author: "Unknown",
		},
	}

	// Create a mock for quote storage.
	storageMock := mocks.NewMockQuoteStorage(quoteList, nil)

	// Create a quote service.
	service := quotes.NewService(zap.NewNop(), storageMock)

	// Call the method under test.
	quote, err := service.GetRandomQuote(context.TODO())

	// Assert no error is returned.
	assert.NoError(t, err)

	// Assert a random quote is returned.
	assert.Contains(t, quoteList, *quote)
}

func TestService_GetRandomQuote_StorageError(t *testing.T) {
	// Prepare test data.
	storageError := errors.New("storage error")

	// Create a mock for quote storage.
	storageMock := mocks.NewMockQuoteStorage(nil, storageError)

	// Create a quote service.
	service := quotes.NewService(zap.NewNop(), storageMock)

	// Call the method under test.
	quote, err := service.GetRandomQuote(context.TODO())

	// Assert the error is returned.
	assert.ErrorIs(t, err, storageError)

	// Assert the quote is nil.
	assert.Nil(t, quote)
}

func TestService_GetRandomQuote_StorageEmpty(t *testing.T) {
	// Prepare test data.
	quoteList := []model.Quote{}

	// Create a mock for quote storage.
	storageMock := mocks.NewMockQuoteStorage(quoteList, nil)

	// Create a quote service.
	service := quotes.NewService(zap.NewNop(), storageMock)

	// Call the method under test.
	quote, err := service.GetRandomQuote(context.TODO())

	// Assert the error is returned.
	assert.Error(t, err, "expected an error")

	// Assert the quote is nil.
	assert.Nil(t, quote)
}
