package integration_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daniel-orlov/quotes-server/internal/domain/model"
	"github.com/daniel-orlov/quotes-server/internal/transport/middleware/proofer"
	"github.com/daniel-orlov/quotes-server/pkg/hashcash"
)

// Happy path test.
func TestIntegration_Server_ReturnsStatus428AndHashcashChallenge_SolveAndReceiveQuote(t *testing.T) {
	// Prepare endpoint
	url := fmt.Sprintf("%s/v1/quotes/random", testServer.URL)

	// Make the request to the server
	resp, err := testClient.Get(url)
	assert.NoError(t, err, "making request to the server failed")
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			t.Logf("closing response body: %v", err)
		}
	}(resp.Body)

	// Assert the initial response status code is 428
	assert.Equal(t, http.StatusPreconditionRequired, resp.StatusCode, "initial response status code is not 428")

	// Retrieve the hashcash challenge from the response header
	hashcashChallengeStr := resp.Header.Get(proofer.ChallengeHeader)
	assert.NotEmpty(t, hashcashChallengeStr, "hashcash challenge is empty")

	// Parse the hashcash challenge
	hashcashChallenge, err := hashcash.ParseStr(hashcashChallengeStr)
	assert.NoError(t, err, "parsing hashcash challenge failed")

	// Solve the hashcash challenge
	solution, err := hashcashChallenge.Solve()
	assert.NoError(t, err, "solving hashcash challenge failed")

	// Create a new request with the solution
	req, err := http.NewRequest(http.MethodGet, url, nil)
	assert.NoError(t, err)
	req.Header.Set(proofer.ChallengeHeader, solution)

	// Make the request with the solution
	resp, err = testClient.Do(req)
	assert.NoError(t, err)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			t.Logf("closing response body: %v", err)
		}
	}(resp.Body)

	// Assert the final response status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read and validate the response body
	// Parse the response body to a Quote
	quote := model.Quote{}

	// Read the response body
	resBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err, "reading response body failed")

	// Unmarshal the response body
	err = json.Unmarshal(resBody, &quote)
	assert.NoError(t, err, "unmarshaling response body failed")

	// Assert the quote is not empty
	assert.NotEmpty(t, quote, "quote is empty")
	assert.NotEmpty(t, quote.Text, "quote text is empty")
	assert.NotEmpty(t, quote.Author, "quote author is empty")
}

func TestIntegration_Server_AttemptsToReuseHashcashSolution_ReturnsError(t *testing.T) {
	// Prepare endpoint
	url := fmt.Sprintf("%s/v1/quotes/random", testServer.URL)

	// Make the initial request to receive the hashcash challenge
	resp, err := testClient.Get(url)
	assert.NoError(t, err)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			t.Logf("closing response body: %v", err)
		}
	}(resp.Body)

	// Assert the initial response status code
	assert.Equal(t, http.StatusPreconditionRequired, resp.StatusCode)

	// Retrieve the hashcash challenge from the response header
	hashcashChallengeStr := resp.Header.Get(proofer.ChallengeHeader)
	assert.NotEmpty(t, hashcashChallengeStr, "hashcash challenge is empty")

	// Parse the hashcash challenge
	hashcashChallenge, err := hashcash.ParseStr(hashcashChallengeStr)
	assert.NoError(t, err, "parsing hashcash challenge failed")

	// Solve the hashcash challenge
	solution, err := hashcashChallenge.Solve()
	assert.NoError(t, err, "solving hashcash challenge failed")

	// Create a new request with the solution
	req, err := http.NewRequest(http.MethodGet, url, nil)
	assert.NoError(t, err)
	req.Header.Set(proofer.ChallengeHeader, solution)

	// Make the request with the solution
	resp, err = testClient.Do(req)
	assert.NoError(t, err)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			t.Logf("closing response body: %v", err)
		}
	}(resp.Body)

	// Assert the final response status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Attempt to reuse the same solution by sending another request with the same solution
	resp, err = testClient.Do(req)
	assert.NoError(t, err)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			t.Logf("closing response body: %v", err)
		}
	}(resp.Body)

	// Assert the response status code
	assert.Equal(t, http.StatusPreconditionRequired, resp.StatusCode)

	// Assert the new challenge is received
	newHashcashChallengeStr := resp.Header.Get(proofer.ChallengeHeader)
	assert.NotEmpty(t, newHashcashChallengeStr, "new hashcash challenge is empty")

	// Assert the new challenge is different from the previous one
	assert.NotEqual(t, hashcashChallengeStr, newHashcashChallengeStr, "new hashcash challenge is the same as the previous one")
}

func TestIntegration_Server_AttemptsToSendWrongChallengeSolution_ReturnsNewChallenge(t *testing.T) {
	// Prepare endpoint
	url := fmt.Sprintf("%s/v1/quotes/random", testServer.URL)

	// Make the initial request to receive the hashcash challenge
	resp, err := testClient.Get(url)
	assert.NoError(t, err)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			t.Logf("closing response body: %v", err)
		}
	}(resp.Body)

	// Assert the initial response status code
	assert.Equal(t, http.StatusPreconditionRequired, resp.StatusCode)

	// Retrieve the hashcash challenge from the response header
	hashcashChallengeStr := resp.Header.Get(proofer.ChallengeHeader)
	assert.NotEmpty(t, hashcashChallengeStr, "hashcash challenge is empty")

	// Ignore the hashcash challenge and send a wrong solution

	// Create a new request with the wrong solution
	wrongSolution := "wrong-solution"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	assert.NoError(t, err)
	req.Header.Set(proofer.ChallengeHeader, wrongSolution)

	// Make the request with the wrong solution
	resp, err = testClient.Do(req)
	assert.NoError(t, err)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			t.Logf("closing response body: %v", err)
		}
	}(resp.Body)

	// Assert the response status code
	assert.Equal(t, http.StatusPreconditionRequired, resp.StatusCode)

	// Assert the new challenge is received
	newHashcashChallengeStr := resp.Header.Get(proofer.ChallengeHeader)
	assert.NotEmpty(t, newHashcashChallengeStr, "new hashcash challenge is empty")

	// Assert the new challenge is different from the previous one
	assert.NotEqual(t, hashcashChallengeStr, newHashcashChallengeStr, "new hashcash challenge is the same as the previous one")
}
