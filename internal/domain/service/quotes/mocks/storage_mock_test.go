package mocks_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
	"github.com/daniel-orlov/quotes-server/internal/domain/service/quotes/mocks"
)

func TestNewMockQuoteStorage_GetQuoteList_NoErrorPassed(t *testing.T) {
	// Prepare test data.
	quotes := []model.Quote{
		{
			ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
			Text:   "Each problem that I solved became a rule which served afterwards to solve other problems.",
			Author: "Rene Descartes",
		},
		{
			ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
			Text:   "The only thing that is constant is change.",
			Author: "Heraclitus",
		},
	}

	// Prepare mock storage, passing nil as an error.
	mockStorage := mocks.NewMockQuoteStorage(quotes, nil)

	// Call the method under test.
	quoteList, err := mockStorage.GetQuoteList(context.TODO())

	// Assert no error is returned.
	assert.NoError(t, err)

	// Assert the quote list is returned.
	assert.Equal(t, quotes, quoteList)
}

func TestNewMockQuoteStorage_GetQuoteList_ErrorPassed(t *testing.T) {
	// Prepare test data.
	storageError := errors.New("storage error")

	// Prepare mock storage, passing an error.
	mockStorage := mocks.NewMockQuoteStorage(nil, storageError)

	// Call the method under test.
	quoteList, err := mockStorage.GetQuoteList(context.TODO())

	// Assert the error is returned.
	assert.Equal(t, storageError, err)

	// Assert the quote list is nil.
	assert.Nil(t, quoteList)
}

func TestNewMockQuoteStorage_GetQuoteList_NoErrorPassed_NoQuotes(t *testing.T) {
	// Prepare mock storage, passing nil as an error and nil as a quote list.
	mockStorage := mocks.NewMockQuoteStorage(nil, nil)

	// Call the method under test.
	quoteList, err := mockStorage.GetQuoteList(context.TODO())

	// Assert no error is returned.
	assert.NoError(t, err)

	// Assert the quote list is nil.
	assert.Nil(t, quoteList)
}
