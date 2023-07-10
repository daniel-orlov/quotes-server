package mocks_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
	"github.com/daniel-orlov/quotes-server/internal/transport/http/quotes/mocks"
)

func TestMockQuoteService_GetRandomQuote_ErrorPassed(t *testing.T) {
	// Add test data
	quote := model.Quote{
		Text:   "Victory is always possible for the person who refuses to stop fighting.",
		Author: "Napoleon Hill",
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
	}

	// Create a mock service
	mockService := mocks.NewMockQuoteService(map[ulid.ULID]model.Quote{
		quote.ID: quote,
	}, errors.New("service error"))

	// Call the method
	_, err := mockService.GetRandomQuote(context.TODO())

	// Assert the error
	assert.Error(t, err)

	// Assert the error message
	assert.Equal(t, "service error", err.Error())
}

func TestMockQuoteService_GetRandomQuote_NoError(t *testing.T) {
	// Add test data
	quote := model.Quote{
		Text:   "It is not the mountain we conquer but ourselves.",
		Author: "Edmund Hillary",
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
	}

	// Create a mock service
	mockService := mocks.NewMockQuoteService(map[ulid.ULID]model.Quote{
		quote.ID: quote,
	}, nil)

	// Call the method
	result, err := mockService.GetRandomQuote(context.TODO())

	// Assert no error
	assert.NoError(t, err)

	// Assert the text of the quote
	assert.Equal(t, quote.Text, result.Text)

	// Assert the author of the quote
	assert.Equal(t, quote.Author, result.Author)

	// Assert the ID of the quote
	assert.Equal(t, quote.ID, result.ID)
}

func TestMockQuoteService_GetQuote_NoQuotesArePassed(t *testing.T) {
	// Create a mock service
	mockService := mocks.NewMockQuoteService(map[ulid.ULID]model.Quote{}, nil)

	// Call the method
	result, err := mockService.GetRandomQuote(context.TODO())

	// Assert the error
	assert.NoError(t, err)

	// Assert the result is nil
	assert.Nil(t, result)
}
