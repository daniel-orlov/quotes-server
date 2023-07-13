package pow_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/pkg/pow"
	"github.com/daniel-orlov/quotes-server/pkg/pow/mocks"
)

func TestService_CheckSolution(t *testing.T) {
	t.Run("Wrong key passed", func(t *testing.T) {
		t.Run("Nil key", func(t *testing.T) {
			// Create a new service
			service := pow.NewService(zap.NewNop(), nil)

			// Check if the solution is correct
			isCorrect, err := service.CheckSolution(context.TODO(), "", nil)

			// Expect an error - ErrChallengeKeyEmpty
			assert.ErrorIs(t, err, pow.ErrChallengeKeyEmpty, "expected ErrChallengeKeyEmpty")

			// Expect the solution to be incorrect
			assert.False(t, isCorrect, "expected false")
		})

		t.Run("Empty key", func(t *testing.T) {
			// Create a new service
			service := pow.NewService(zap.NewNop(), nil)

			// Check if the solution is correct
			isCorrect, err := service.CheckSolution(context.TODO(), "", &pow.ChallengeKey{})

			// Expect an error - ErrChallengeKeyEmpty
			assert.ErrorIs(t, err, pow.ErrChallengeKeyEmpty, "expected ErrChallengeKeyEmpty")

			// Expect the solution to be incorrect
			assert.False(t, isCorrect, "expected false")
		})
	})

	t.Run("Storage errors", func(t *testing.T) {
		t.Run("Solution does not exist in the store", func(t *testing.T) {
			// Create mock storage
			store := mocks.NewMockChallengeStorage(nil, nil)

			// Create a new service
			service := pow.NewService(zap.NewNop(), store)

			// Check if the solution is correct
			isCorrect, err := service.CheckSolution(context.TODO(), "", pow.NewChallengeKey("clientID", "resourceID"))

			// Expect an error - ErrChallengeNotFound
			assert.ErrorIs(t, err, pow.ErrChallengeNotFound, "expected ErrChallengeNotFound")

			// Expect the solution to be incorrect
			assert.False(t, isCorrect, "expected false")
		})

		t.Run("Error getting challenge from the store", func(t *testing.T) {
			// Create mock storage
			store := mocks.NewMockChallengeStorage(nil, errors.New("error"))

			// Create a new service
			service := pow.NewService(zap.NewNop(), store)

			// Check if the solution is correct
			isCorrect, err := service.CheckSolution(context.TODO(), "", pow.NewChallengeKey("clientID", "resourceID"))

			// Expect an error
			assert.Error(t, err, "expected error")

			// Expect the solution
			assert.False(t, isCorrect, "expected false")
		})
	})

	t.Run("Solution errors", func(t *testing.T) {
		t.Run("Solution is incorrect", func(t *testing.T) {
			// Create mock storage
			store := mocks.NewMockChallengeStorage(
				map[string]string{
					"clientID:resourceID": "1:20:23:not-solved::Kl7oUEQg:4c73d",
				}, nil)

			// Create a new service
			service := pow.NewService(zap.NewNop(), store)

			// Check if the solution is correct
			isCorrect, err := service.CheckSolution(context.TODO(), "1:20:23:not-solved::Kl7oUEQg:4c73d", pow.NewChallengeKey("clientID", "resourceID"))

			// Expect an error
			assert.Error(t, err, "expected error")

			// Expect the solution
			assert.False(t, isCorrect, "expected false")
		})

		t.Run("Solution is invalid", func(t *testing.T) {
			// Create mock storage
			store := mocks.NewMockChallengeStorage(map[string]string{
				"clientID:resourceID": "invalid",
			}, nil)

			// Create a new service
			service := pow.NewService(zap.NewNop(), store)

			// Check if the solution is correct
			isCorrect, err := service.CheckSolution(context.TODO(), "invalid", pow.NewChallengeKey("clientID", "resourceID"))

			// Expect an error
			assert.Error(t, err, "expected error")

			// Expect the solution
			assert.False(t, isCorrect, "expected false")
		})
	})

	t.Run("Solution is correct", func(t *testing.T) {
		t.Run("Solution is correct", func(t *testing.T) {
			// Create mock storage
			store := mocks.NewMockChallengeStorage(
				map[string]string{
					"clientID:resourceID": "1:20:23:some-resource::Kl7oUEQg:0",
				}, nil)

			// Create a new service
			service := pow.NewService(zap.NewNop(), store)

			// Check if the solution is correct
			isCorrect, err := service.CheckSolution(context.TODO(), "1:20:23:some-resource::Kl7oUEQg:4c73d", pow.NewChallengeKey("clientID", "resourceID"))

			// Expect no error
			assert.NoError(t, err, "expected no error")

			// Expect the solution
			assert.True(t, isCorrect, "expected true")
		})
	})
}
