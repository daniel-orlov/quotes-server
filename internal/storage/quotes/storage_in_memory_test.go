package quotes_test

import (
	"context"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
	"github.com/daniel-orlov/quotes-server/internal/storage/quotes"
)

func TestStorageInMemory_GetQuoteList_Success(t *testing.T) {
	// Prepare test data.
	quoteList := []model.Quote{
		{
			ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
			Text:   "People will forget what you said, people will forget what you did, but people will never forget how you made them feel.",
			Author: "Maya Angelou",
		},
		{
			ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
			Text:   "The best way to cheer yourself up is to try to cheer somebody else up.",
			Author: "Mark Twain",
		},
	}

	// Create a quote storage.
	storage := quotes.NewStorageInMemory(zap.NewNop(), quoteList)

	// Call the method under test.
	quoteListActual, err := storage.GetQuoteList(context.TODO())

	// Assert no error is returned.
	assert.NoError(t, err)

	// Assert the quote list is returned.
	assert.Equal(t, quoteList, quoteListActual)
}

func TestStorageInMemory_GetQuoteList_Empty(t *testing.T) {
	// Prepare test data.
	quoteList := []model.Quote{}

	// Create a quote storage.
	storage := quotes.NewStorageInMemory(zap.NewNop(), quoteList)

	// Call the method under test.
	quoteListActual, err := storage.GetQuoteList(context.TODO())

	// Assert error is returned.
	assert.ErrorIs(t, err, quotes.ErrDBEmpty)

	// Assert an empty quote list is returned.
	assert.Empty(t, quoteListActual)
}
