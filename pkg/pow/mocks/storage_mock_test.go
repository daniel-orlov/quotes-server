package mocks_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daniel-orlov/quotes-server/pkg/pow/mocks"
)

func TestMockChallengeStorage_Add(t *testing.T) {
	t.Run("Error was set", func(t *testing.T) {
		// Create a new mock store with an error
		store := mocks.NewMockChallengeStorage(nil, errors.New("error"))

		// Add a challenge to the store
		err := store.Add(context.TODO(), "key", "value")

		// Check if the error is correct
		assert.Error(t, err, "expected error")
	})

	t.Run("Error was not set", func(t *testing.T) {
		// Create a new mock store
		store := mocks.NewMockChallengeStorage(nil, nil)

		// Add a challenge to the store
		err := store.Add(context.TODO(), "key", "value")

		// Check if the error is correct
		assert.NoError(t, err, "expected no error")
	})
}

func TestMockChallengeStorage_Get(t *testing.T) {
	t.Run("Error was set", func(t *testing.T) {
		// Create a new mock store with an error
		store := mocks.NewMockChallengeStorage(nil, errors.New("error"))

		// Get a challenge from the store
		exists, err := store.Get(context.TODO(), "key")

		// Check if the error is correct
		assert.Error(t, err, "expected error")

		// Check if the result is correct
		assert.False(t, exists, "expected false")
	})

	t.Run("Error was not set", func(t *testing.T) {
		t.Run("Challenge exists", func(t *testing.T) {
			// Create a new mock store
			store := mocks.NewMockChallengeStorage(map[string]string{"key": "value"}, nil)

			// Get a challenge from the store
			exists, err := store.Get(context.TODO(), "key")

			// Check if the error is correct
			assert.NoError(t, err, "expected no error")

			// Check if the result is correct
			assert.True(t, exists, "expected true")
		})

		t.Run("Challenge does not exist", func(t *testing.T) {
			// Create a new mock store
			store := mocks.NewMockChallengeStorage(map[string]string{"key": "value"}, nil)

			// Get a challenge from the store
			exists, err := store.Get(context.TODO(), "key2")

			// Check if the error is correct
			assert.NoError(t, err, "expected no error")

			// Check if the result is correct
			assert.False(t, exists, "expected false")
		})
	})
}

func TestMockChallengeStorage_Delete(t *testing.T) {
	t.Run("Error was set", func(t *testing.T) {
		// Create a new mock store with an error
		store := mocks.NewMockChallengeStorage(nil, errors.New("error"))

		// Delete a challenge from the store
		err := store.Delete(context.TODO(), "key")

		// Check if the error is correct
		assert.Error(t, err, "expected error")
	})

	t.Run("Error was not set", func(t *testing.T) {
		// Create a new mock store
		store := mocks.NewMockChallengeStorage(nil, nil)

		// Delete a challenge from the store
		err := store.Delete(context.TODO(), "key")

		// Check if the error is correct
		assert.NoError(t, err, "expected no error")
	})
}
