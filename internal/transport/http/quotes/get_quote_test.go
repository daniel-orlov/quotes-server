package quotes_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
	"github.com/daniel-orlov/quotes-server/internal/transport/http/quotes"
	"github.com/daniel-orlov/quotes-server/internal/transport/http/quotes/mocks"
)

func TestHandler_GetQuote_Success(t *testing.T) {
	// Setting the gin to test mode
	gin.SetMode(gin.TestMode)
	// Creating a recorder to record the response
	w := httptest.NewRecorder()
	// Creating a context to use in the request
	c, r := gin.CreateTestContext(w)

	// Creating test data
	quoteInService := model.Quote{
		Text:   "The best preparation for tomorrow is doing your best today.",
		Author: "H. Jackson Brown, Jr.",
		ID:     ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy()),
	}

	// Creating a mock service
	quoteService := mocks.NewMockQuoteService(map[ulid.ULID]model.Quote{
		quoteInService.ID: quoteInService,
	},
		nil,
	)

	// Creating a handler
	quoteHandler := quotes.NewHandler(quoteService, zap.NewNop())

	// Creating a variable to store the endpoint of the request
	endpoint := "/v1/quotes"

	// Registering the endpoint handler to the router
	r.GET(endpoint, quoteHandler.GetQuote)

	// Creating a request
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)

	// Requiring no error
	require.NoError(t, err)

	// Serving the request
	r.ServeHTTP(c.Writer, req)

	// Asserting the response code
	assert.Equal(t, http.StatusOK, w.Code)

	// Creating a variable to store the response body
	var quote model.Quote

	// Unmarshalling the response body into the quote variable
	err = json.Unmarshal(w.Body.Bytes(), &quote)

	// Asserting the error is nil
	assert.NoError(t, err)

	// Asserting the quote is not nil
	assert.NotNil(t, quote)

	// Asserting the quote text is not empty
	assert.NotEmpty(t, quote.Text)

	// Asserting the quote author is not empty
	assert.NotEmpty(t, quote.Author)

	// Asserting the quote text is correct
	assert.Equal(t, quoteInService.Text, quote.Text)

	// Asserting the quote author is correct
	assert.Equal(t, quoteInService.Author, quote.Author)

	// Asserting the quote id is not empty
	assert.NotEmpty(t, quote.ID)

	// Asserting the quote id is correct
	assert.Equal(t, quoteInService.ID, quote.ID)
}

func TestHandler_GetQuote_Internal_Error(t *testing.T) {
	// Setting the gin to test mode
	gin.SetMode(gin.TestMode)
	// Creating a recorder to record the response
	w := httptest.NewRecorder()
	// Creating a context to use in the request
	c, r := gin.CreateTestContext(w)

	// Creating a mock service
	quoteService := mocks.NewMockQuoteService(nil, errors.New("internal error"))

	// Creating a handler
	quoteHandler := quotes.NewHandler(quoteService, zap.NewNop())

	// Creating a variable to store the endpoint of the request
	endpoint := "/v1/quotes"

	// Registering the endpoint handler to the router
	r.GET(endpoint, quoteHandler.GetQuote)

	// Creating a request
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)

	// Requiring no error
	require.NoError(t, err)

	// Serving the request
	r.ServeHTTP(c.Writer, req)

	// Asserting the response code
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Creating a variable to store the response body
	var quote model.Quote

	// Unmarshalling the response body into the quote variable
	err = json.Unmarshal(w.Body.Bytes(), &quote)

	// Asserting the error is nil
	assert.NoError(t, err)

	// Asserting the quote is not nil
	assert.NotNil(t, quote)

	// Asserting the quote text is empty
	assert.Empty(t, quote.Text)

	// Asserting the quote author is empty
	assert.Empty(t, quote.Author)

	// Asserting the quote id is empty
	assert.Empty(t, quote.ID)
}
