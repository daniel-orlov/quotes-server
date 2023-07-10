package http_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
	httptransport "github.com/daniel-orlov/quotes-server/internal/transport/http" //
	"github.com/daniel-orlov/quotes-server/internal/transport/http/quotes"
	"github.com/daniel-orlov/quotes-server/internal/transport/http/quotes/mocks"
)

func TestNewRouter_GET_quote(t *testing.T) {
	// Prepare test data
	quoteInService := model.Quote{
		Text:   "We are what we repeatedly do. Excellence, then, is not an act, but a habit.",
		Author: "Aristotle",
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
	}

	// Create a handler
	quoteHandler := quotes.NewHandler(mocks.NewMockQuoteService(map[ulid.ULID]model.Quote{
		quoteInService.ID: quoteInService,
	}, nil), zap.NewNop())

	// Create a router
	r := httptransport.NewRouter(quoteHandler)

	// Create a recorder to record the response
	w := httptest.NewRecorder()

	// Create a request
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1%s", quotes.ResourceEndpoint), nil)

	// Require no error
	require.NoError(t, err)

	// Serve the request
	r.ServeHTTP(w, req)

	// Assert the response code
	assert.Equal(t, http.StatusOK, w.Code)
}
