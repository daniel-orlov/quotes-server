package proofer_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/internal/transport/middleware/proofer"
	"github.com/daniel-orlov/quotes-server/internal/transport/middleware/proofer/mocks"
)

const testEndpoint = "/test"

func TestProofer_Use(t *testing.T) {
	t.Run("Challenge Request", func(t *testing.T) {
		t.Run("Service failure", func(t *testing.T) {
			// Create a mock PoW service
			svc := mocks.NewMockPoWService("", false, errors.New("service failure"))

			// Create a Proofer instance with mock dependencies
			mw := proofer.New(zap.NewNop(), &proofer.Config{}, svc)

			// Setting the gin to test mode
			gin.SetMode(gin.TestMode)
			// Creating a recorder to record the response
			w := httptest.NewRecorder()
			// Creating a context to use in the request
			c, r := gin.CreateTestContext(w)

			// Create a Gin handler using the Proofer middleware
			r.GET(testEndpoint, mw.Use())

			// Creating a request
			req := httptest.NewRequest(http.MethodGet, testEndpoint, nil)

			// Serving the request
			r.ServeHTTP(c.Writer, req)

			// Assertions
			assert.Equal(t, http.StatusInternalServerError, w.Code, "status code should be 500")
			assert.Equal(t, "{\"error\":\"failed to get new challenge\"}", w.Body.String())
		})

		t.Run("Get back a challenge", func(t *testing.T) {
			// Create a mock PoW service
			svc := mocks.NewMockPoWService("challenge", false, nil)

			// Create a Proofer instance with mock dependencies
			mw := proofer.New(zap.NewNop(), &proofer.Config{}, svc)

			// Setting the gin to test mode
			gin.SetMode(gin.TestMode)
			// Creating a recorder to record the response
			w := httptest.NewRecorder()
			// Creating a context to use in the request
			c, r := gin.CreateTestContext(w)

			// Create a Gin handler using the Proofer middleware
			r.GET(testEndpoint, mw.Use())

			// Creating a request
			req := httptest.NewRequest(http.MethodGet, testEndpoint, nil)

			// Serving the request
			r.ServeHTTP(c.Writer, req)

			// Assertions
			assert.Equal(t, http.StatusPreconditionRequired, w.Code, "status code should be 428")
			assert.Equal(t, "challenge", w.Header().Get(proofer.ChallengeHeader))
		})

		t.Run("Checking solution", func(t *testing.T) {
			t.Run("Service failure", func(t *testing.T) {
				// Create a mock PoW service
				svc := mocks.NewMockPoWService("", false, errors.New("service failure"))

				// Create a Proofer instance with mock dependencies
				mw := proofer.New(zap.NewNop(), &proofer.Config{}, svc)

				// Setting the gin to test mode
				gin.SetMode(gin.TestMode)
				// Creating a recorder to record the response
				w := httptest.NewRecorder()
				// Creating a context to use in the request
				c, r := gin.CreateTestContext(w)

				// Create a Gin handler using the Proofer middleware
				r.GET(testEndpoint, mw.Use())

				// Creating a request
				req := httptest.NewRequest(http.MethodGet, testEndpoint, nil)
				req.Header.Set(proofer.ChallengeHeader, "challenge")

				// Serving the request
				r.ServeHTTP(c.Writer, req)

				// Assertions
				assert.Equal(t, http.StatusInternalServerError, w.Code, "status code should be 500")
				assert.Equal(t, "{\"error\":\"failed to get new challenge\"}{\"error\":\"failed to check solution\"}", w.Body.String())
			})

			t.Run("Incorrect solution, get back with a new one", func(t *testing.T) {
				// Create a mock PoW service
				svc := mocks.NewMockPoWService("new_challenge", false, nil)

				// Create a Proofer instance with mock dependencies
				mw := proofer.New(zap.NewNop(), &proofer.Config{}, svc)

				// Setting the gin to test mode
				gin.SetMode(gin.TestMode)
				// Creating a recorder to record the response
				w := httptest.NewRecorder()
				// Creating a context to use in the request
				c, r := gin.CreateTestContext(w)

				// Create a Gin handler using the Proofer middleware
				r.GET(testEndpoint, mw.Use())

				// Creating a request
				req := httptest.NewRequest(http.MethodGet, testEndpoint, nil)
				req.Header.Set(proofer.ChallengeHeader, "solution")

				// Serving the request
				r.ServeHTTP(c.Writer, req)

				// Assertions
				assert.Equal(t, http.StatusPreconditionRequired, w.Code, "status code should be 428")
				assert.Equal(t, "new_challenge", w.Header().Get(proofer.ChallengeHeader))
			})

			t.Run("Correct solution", func(t *testing.T) {
				// Create a mock PoW service
				svc := mocks.NewMockPoWService("", true, nil)

				// Create a Proofer instance with mock dependencies
				mw := proofer.New(zap.NewNop(), &proofer.Config{}, svc)

				// Setting the gin to test mode
				gin.SetMode(gin.TestMode)
				// Creating a recorder to record the response
				w := httptest.NewRecorder()
				// Creating a context to use in the request
				c, r := gin.CreateTestContext(w)

				// Create a Gin handler using the Proofer middleware
				r.GET(testEndpoint, mw.Use())

				// Creating a request
				req := httptest.NewRequest(http.MethodGet, testEndpoint, nil)
				req.Header.Set(proofer.ChallengeHeader, "solution")

				// Serving the request
				r.ServeHTTP(c.Writer, req)

				// Assertions
				assert.Equal(t, http.StatusOK, w.Code, "status code should be 200")
			})
		})
	})
}
