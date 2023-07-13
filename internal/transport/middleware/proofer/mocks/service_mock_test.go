package mocks_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daniel-orlov/quotes-server/internal/transport/middleware/proofer/mocks"
)

func TestMockPoWService_CheckSolution(t *testing.T) {
	t.Run("Service error", func(t *testing.T) {
		// Create a new service
		service := mocks.NewMockPoWService("", false, errors.New("service error"))

		// Check the solution
		isCorrect, err := service.CheckSolution(context.TODO(), "", nil)

		// Expect an error
		assert.Error(t, err, "expected an error")

		// Expect the solution to be incorrect
		assert.False(t, isCorrect, "expected false")
	})

	t.Run("Solution is incorrect", func(t *testing.T) {
		// Create a new service
		service := mocks.NewMockPoWService("", false, nil)

		// Check the solution
		isCorrect, err := service.CheckSolution(context.TODO(), "", nil)

		// Expect no error
		assert.NoError(t, err, "expected no error")

		// Expect the solution to be incorrect
		assert.False(t, isCorrect, "expected false")
	})

	t.Run("Solution is correct", func(t *testing.T) {
		// Create a new service
		service := mocks.NewMockPoWService("", true, nil)

		// Check the solution
		isCorrect, err := service.CheckSolution(context.TODO(), "", nil)

		// Expect no error
		assert.NoError(t, err, "expected no error")

		// Expect the solution to be correct
		assert.True(t, isCorrect, "expected true")
	})
}

func TestMockPoWService_NewChallenge(t *testing.T) {
	t.Run("Service error", func(t *testing.T) {
		// Create a new service
		service := mocks.NewMockPoWService("", false, errors.New("service error"))

		// Create a new challenge
		challenge, err := service.NewChallenge(context.Background(), nil, 0, 0)

		// Expect an error
		assert.Error(t, err, "expected an error")

		// Expect the challenge to be empty
		assert.Empty(t, challenge, "expected empty")
	})

	t.Run("Success", func(t *testing.T) {
		// Prepare the challenge
		challenge := "new-challenge"

		// Create a new service
		service := mocks.NewMockPoWService("new-challenge", false, nil)

		// Create a new challenge
		newChallenge, err := service.NewChallenge(context.Background(), nil, 0, 0)

		// Expect no error
		assert.NoError(t, err, "expected no error")

		// Expect a new challenge to be equal to the prepared one
		assert.Equal(t, challenge, newChallenge, "expected equal")
	})
}
