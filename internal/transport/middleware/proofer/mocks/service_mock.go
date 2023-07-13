package mocks

import (
	"context"

	"github.com/daniel-orlov/quotes-server/pkg/pow"
)

// MockPoWService is a mock implementation of the PoW service.
type MockPoWService struct {
	challenge       string
	challengeSolved bool
	serviceError    error
}

// NewMockPoWService creates a new mock PoW service.
func NewMockPoWService(challenge string, challengeSolved bool, serviceError error) *MockPoWService {
	return &MockPoWService{challenge: challenge, challengeSolved: challengeSolved, serviceError: serviceError}
}

// NewChallenge generates a new challenge.
func (m MockPoWService) NewChallenge(_ context.Context, _ pow.Key, _, _ int) (string, error) {
	// If the service error is not nil, return it
	if m.serviceError != nil {
		return "", m.serviceError
	}

	// Return the challenge
	return m.challenge, nil
}

// CheckSolution checks if the solution is valid.
func (m MockPoWService) CheckSolution(_ context.Context, _ string, _ pow.Key) (bool, error) {
	// If the service error is not nil, return it
	if m.serviceError != nil {
		return false, m.serviceError
	}

	// Return the result
	return m.challengeSolved, nil
}
