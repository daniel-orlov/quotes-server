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

func TestService_NewChallenge(t *testing.T) {
	t.Run("Challenge Key is empty", func(t *testing.T) {
		t.Run("Nil key", func(t *testing.T) {
			// Create a new service
			service := pow.NewService(zap.NewNop(), nil)

			// Generate a new challenge
			challenge, err := service.NewChallenge(context.TODO(), nil, 20, 8)

			// Expect an error - ErrChallengeKeyEmpty
			assert.ErrorIs(t, err, pow.ErrChallengeKeyEmpty, "expected ErrChallengeKeyEmpty")

			// Expect the challenge to be empty
			assert.Empty(t, challenge, "expected empty challenge")
		})

		t.Run("Empty key", func(t *testing.T) {
			// Create a new service
			service := pow.NewService(zap.NewNop(), nil)

			// Generate a new challenge
			challenge, err := service.NewChallenge(context.TODO(), &pow.ChallengeKey{}, 20, 8)

			// Expect an error - ErrChallengeKeyEmpty
			assert.ErrorIs(t, err, pow.ErrChallengeKeyEmpty, "expected ErrChallengeKeyEmpty")

			// Expect the challenge to be empty
			assert.Empty(t, challenge, "expected empty challenge")
		})

		t.Run("Empty clientID", func(t *testing.T) {
			// Create a new service
			service := pow.NewService(zap.NewNop(), nil)

			// Generate a new challenge
			challenge, err := service.NewChallenge(context.TODO(), pow.NewChallengeKey("", "resourceID"), 20, 8)

			// Expect an error - ErrChallengeKeyEmpty
			assert.ErrorIs(t, err, pow.ErrChallengeKeyEmpty, "expected ErrChallengeKeyEmpty")

			// Expect the challenge to be empty
			assert.Empty(t, challenge, "expected empty challenge")
		})

		t.Run("Empty resourceID", func(t *testing.T) {
			// Create a new service
			service := pow.NewService(zap.NewNop(), nil)

			// Generate a new challenge
			challenge, err := service.NewChallenge(context.TODO(), pow.NewChallengeKey("clientID", ""), 20, 8)

			// Expect an error - ErrChallengeKeyEmpty
			assert.ErrorIs(t, err, pow.ErrChallengeKeyEmpty, "expected ErrChallengeKeyEmpty")

			// Expect the challenge to be empty
			assert.Empty(t, challenge, "expected empty challenge")
		})
	})

	t.Run("Difficulty is invalid", func(t *testing.T) {
		// Create a new service
		service := pow.NewService(zap.NewNop(), nil)

		// Generate a new challenge
		challenge, err := service.NewChallenge(context.TODO(), pow.NewChallengeKey("clientID", "resourceID"), -1, 8)

		// Expect an error - ErrChallengeDifficultyInvalid
		assert.ErrorIs(t, err, pow.ErrChallengeDifficultyInvalid, "expected ErrChallengeDifficultyInvalid")

		// Expect the challenge to be empty
		assert.Empty(t, challenge, "expected empty challenge")
	})

	t.Run("Salt length is invalid", func(t *testing.T) {
		// Create a new service
		service := pow.NewService(zap.NewNop(), nil)

		// Generate a new challenge
		challenge, err := service.NewChallenge(context.TODO(), pow.NewChallengeKey("clientID", "resourceID"), 20, -1)

		// Expect an error - ErrChallengeSaltLengthInvalid
		assert.ErrorIs(t, err, pow.ErrChallengeSaltLengthInvalid, "expected ErrChallengeSaltLengthInvalid")

		// Expect the challenge to be empty
		assert.Empty(t, challenge, "expected empty challenge")
	})

	t.Run("Storage error", func(t *testing.T) {
		// Create mock storage
		store := mocks.NewMockChallengeStorage(nil, errors.New("storage error"))

		// Create a new service
		service := pow.NewService(zap.NewNop(), store)

		// Generate a new challenge
		challenge, err := service.NewChallenge(context.TODO(), pow.NewChallengeKey("clientID", "resourceID"), 20, 8)

		// Expect an error - storage error
		assert.Error(t, err, "expected storage error")

		// Expect the challenge to be empty
		assert.Empty(t, challenge, "expected empty challenge")
	})

	t.Run("Success", func(t *testing.T) {
		// Create mock storage
		store := mocks.NewMockChallengeStorage(nil, nil)

		// Create a new service
		service := pow.NewService(zap.NewNop(), store)

		// Generate a new challenge
		challenge, err := service.NewChallenge(context.TODO(), pow.NewChallengeKey("clientID", "resourceID"), 20, 8)

		// Expect no error
		assert.NoError(t, err, "expected no error")

		// Expect the challenge to be not empty
		assert.NotEmpty(t, challenge, "expected not empty challenge")
	})
}
