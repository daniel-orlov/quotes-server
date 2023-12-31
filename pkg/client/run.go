package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
	"github.com/daniel-orlov/quotes-server/internal/transport/middleware/proofer"
	"github.com/daniel-orlov/quotes-server/pkg/hashcash"
)

// Run runs the client.
func (c *Client) Run() {
	// Create the endpoint URL
	endpoint := c.buildEndpointURL()

	// Create a function to send a request
	sendRequest := func() {
		if err := c.SendRequest(endpoint); err != nil {
			c.logger.Error("sending request", zap.Error(err), zap.String("endpoint", endpoint))
		}
	}

	// Create a function to run the client in a loop
	runClient := func() {
		// Send a request
		sendRequest()

		// Sleep to control the request rate
		sleepDuration := time.Second / time.Duration(c.cfg.Connection.RequestRatePerSecond)
		// Log the sleep duration
		c.logger.Debug("sleeping", zap.Float64("seconds", sleepDuration.Seconds()))
		// Sleep
		time.Sleep(sleepDuration)
	}

	// If the request count is set, send that many requests
	if c.cfg.Connection.RequestCount > 0 {
		for i := 0; i < c.cfg.Connection.RequestCount; i++ {
			runClient()
		}
	} else {
		for {
			// Otherwise, run the client in a loop indefinitely
			runClient()
		}
	}
}

// SendRequest sends a request to the endpoint.
func (c *Client) SendRequest(url string) error {
	// Create a new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	// Return any error
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	// Send the request
	res, err := c.client.Do(req) //nolint:bodyclose // The response body is closed in defer func and linter can't see it
	// Return any error
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}

	// Close the response body
	defer func(Body io.ReadCloser, logger *zap.Logger) {
		// Close the response body and log any error
		err = Body.Close()
		if err != nil {
			logger.Warn("closing response body", zap.Error(err))
		}
	}(res.Body, c.logger)

	// Log the response
	err = c.logResponse(res)
	if err != nil {
		return fmt.Errorf("logging response: %w", err)
	}

	// Check if the response is a hashcash challenge
	hashcashChallenge := res.Header.Get(proofer.ChallengeHeader)

	// If it does, solve it and retry
	if hashcashChallenge != "" {
		return c.solveHashcashChallengeAndRetry(url, hashcashChallenge)
	}

	return nil
}

// buildEndpointURL builds the endpoint URL.
func (c *Client) buildEndpointURL() string {
	return fmt.Sprintf(
		"http://%s:%d%s",
		c.cfg.Connection.ServerHost,
		c.cfg.Connection.ServerPort,
		c.cfg.Connection.RequestPath,
	)
}

// solveHashcashChallengeAndRetry solves the hashcash challenge and retries the request.
func (c *Client) solveHashcashChallengeAndRetry(url, hashcashChallenge string) error {
	// Parse the hashcash challenge
	hc, err := hashcash.ParseStr(hashcashChallenge)
	// Return any error
	if err != nil {
		return fmt.Errorf("parsing hashcash challenge: %w", err)
	}

	// Solve the hashcash challenge
	solution, err := hc.Solve()
	// Return any error
	if err != nil {
		return fmt.Errorf("solving hashcash challenge: %w", err)
	}

	// Retry the request with the solution
	// Create a new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	// Return any error
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	// Set the solution as a header
	req.Header.Set(proofer.ChallengeHeader, solution)

	// Send the request
	res, err := c.client.Do(req) //nolint:bodyclose // The response body is closed in defer func and linter can't see it
	// Return any error
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}

	// Close the response body
	defer func(Body io.ReadCloser, logger *zap.Logger) {
		// Close the response body and log any error
		err = Body.Close()
		if err != nil {
			logger.Warn("closing response body", zap.Error(err))
		}
	}(res.Body, c.logger)

	// Log the response
	err = c.logResponse(res)
	// Return any error
	if err != nil {
		return fmt.Errorf("logging response: %w", err)
	}

	// Return no error
	return nil
}

// logResponse logs the response.
func (c *Client) logResponse(res *http.Response) error {
	// Parse the response body to a Quote
	quote := model.Quote{}

	// Read the response body
	resBody, err := io.ReadAll(res.Body)
	// Return any error
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	// Unmarshal the response body
	err = json.Unmarshal(resBody, &quote)
	// Return any error
	if err != nil {
		return fmt.Errorf("decoding response body: %w", err)
	}

	// Log the response
	c.logger.Debug("response",
		zap.String("status", res.Status),
		zap.Any("headers", res.Header),
		zap.Any("body", quote),
	)

	// Return no error
	return nil
}
