package pow

import (
	"context"
	"fmt"

	"github.com/daniel-orlov/quotes-server/pkg/hashcash"
)

// NewChallenge generates a new challenge, saves it to the store and returns it.
func (s *Service) NewChallenge(ctx context.Context, key Key, difficulty, saltLength int) (string, error) {
	// Check if the challenge key is empty
	if key == nil || key.String() == "" || key.ResourceID() == "" || key.ClientID() == "" {
		// Return an error if the challenge key is empty
		return "", ErrChallengeKeyEmpty
	}

	// Check if the difficulty is valid
	if difficulty <= 0 {
		// Return an error if the difficulty is invalid
		return "", ErrChallengeDifficultyInvalid
	}

	// Check if the salt length is valid
	if saltLength <= 0 {
		// Return an error if the salt length is invalid
		return "", ErrChallengeSaltLengthInvalid
	}

	// Generate a new challenge
	challenge, err := hashcash.New(difficulty, saltLength, hashcash.DateFormatYYMMDD, key.ClientID())
	if err != nil {
		return "", fmt.Errorf("generating new challenge: %w", err)
	}

	// Stringify the challenge
	challengeStr := challenge.String()

	// Save the challenge to the store
	err = s.Store.Add(ctx, key.String(), challengeStr)
	if err != nil {
		return "", fmt.Errorf("saving challenge to store: %w", err)
	}

	// Return the challenge
	return challengeStr, nil
}
