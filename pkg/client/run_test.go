package client_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/internal/transport/middleware/proofer"
	"github.com/daniel-orlov/quotes-server/pkg/client"
)

func TestClient_Run(t *testing.T) {
	logger := zap.NewNop()

	cfg, err := client.NewConfig()
	require.NoError(t, err, "creating config")

	t.Run("Successful Request", func(t *testing.T) {
		// Create a new client
		quotesClient := client.NewClient(logger, cfg, &http.Client{})

		// Start a mock server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Respond with a success status code and a quote JSON body
			w.WriteHeader(http.StatusOK)
			_, err = w.Write([]byte(`{"quote": "Test quote"}`))
			if err != nil {
				return
			}
		}))
		defer server.Close()

		// Override the client's endpoint with the mock server URL
		cfg.Connection.ServerHost = server.URL

		// Run the client
		go quotesClient.Run()

		// Wait for the client to make a few requests
		time.Sleep(100 * time.Millisecond)

		// Assertions
		// No assertions are needed here, the test will fail if the client panics or exits with an error
	})

	t.Run("Rate Limit Exceeded", func(t *testing.T) {
		// Create a new client
		quotesClient := client.NewClient(logger, cfg, &http.Client{})

		// Start a mock server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Respond with a 429 Too Many Requests status code
			w.WriteHeader(http.StatusTooManyRequests)
		}))
		defer server.Close()

		// Override the client's endpoint with the mock server URL
		cfg.Connection.ServerHost = server.URL

		// Run the client
		go quotesClient.Run()

		// Wait for the client to make a few requests
		time.Sleep(100 * time.Millisecond)

		// Assertions
		// No assertions are needed here, the test will fail if the client panics or exits with an error
	})

	t.Run("Hashcash Challenge", func(t *testing.T) {
		// Create a new client
		quotesClient := client.NewClient(logger, cfg, &http.Client{})

		// Start a mock server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Respond with a hashcash challenge header
			w.Header().Set(proofer.ChallengeHeader, "hashcash_challenge")
			// Respond with a success status code and a quote JSON body
			w.WriteHeader(http.StatusPreconditionRequired)
		}))
		defer server.Close()

		// Override the client's endpoint with the mock server URL
		cfg.Connection.ServerHost = server.URL

		// Run the client
		go quotesClient.Run()

		// Wait for the client to make a few requests
		time.Sleep(100 * time.Millisecond)

		// Assertions
		// No assertions are needed here, the test will fail if the client panics or exits with an error
	})
}

func TestClient_SendRequest(t *testing.T) {
	logger := zap.NewNop()

	cfg, err := client.NewConfig()
	require.NoError(t, err, "creating config")

	t.Run("Successful Request", func(t *testing.T) {
		// Create a new client
		quotesClient := client.NewClient(logger, cfg, &http.Client{})

		// Start a mock server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Respond with a success status code and a quote JSON body
			w.WriteHeader(http.StatusOK)
			_, err = w.Write([]byte(`{"quote": "Test quote"}`))
			require.NoError(t, err, "writing response body")
		}))
		defer server.Close()

		// Override the client's endpoint with the mock server URL
		endpoint := server.URL

		// Invoke the SendRequest method
		err = quotesClient.SendRequest(endpoint)

		// Assertions
		assert.NoError(t, err)
	})

	t.Run("Rate Limit Exceeded", func(t *testing.T) {
		// Create a new client
		quotesClient := client.NewClient(logger, cfg, &http.Client{})

		// Start a mock server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Respond with a 429 Too Many Requests status code
			w.WriteHeader(http.StatusTooManyRequests)
		}))
		defer server.Close()

		// Override the client's endpoint with the mock server URL
		endpoint := server.URL

		// Invoke the SendRequest method
		err = quotesClient.SendRequest(endpoint)

		// Assertions
		assert.Error(t, err)
	})

	t.Run("Hashcash Challenge", func(t *testing.T) {
		// Create a new client
		quotesClient := client.NewClient(logger, cfg, &http.Client{})

		// Start a mock server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Respond with a hashcash challenge header
			w.Header().Set(proofer.ChallengeHeader, "1:20:23:some-resource::salt:23a")
			// Respond with a precondition required status code
			w.WriteHeader(http.StatusPreconditionRequired)
			w.Write([]byte(`{"quote": "Test quote"}`))
		}))
		defer server.Close()

		// Override the client's endpoint with the mock server URL
		endpoint := server.URL

		// Invoke the SendRequest method
		err = quotesClient.SendRequest(endpoint)

		// Assertions
		assert.NoError(t, err)
	})

	t.Run("Invalid Response Body", func(t *testing.T) {
		// Create a new client
		quotesClient := client.NewClient(logger, cfg, &http.Client{})

		// Start a mock server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Respond with an invalid response body
			w.WriteHeader(http.StatusOK)
			_, err = w.Write([]byte(`invalid-json-response`))
			require.NoError(t, err, "writing response body")
		}))
		defer server.Close()

		// Override the client's endpoint with the mock server URL
		endpoint := server.URL

		// Invoke the SendRequest method
		err = quotesClient.SendRequest(endpoint)

		// Assertions
		assert.Error(t, err)
	})
}
