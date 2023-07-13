package pow

import "errors"

var (
	// ErrChallengeNotFound is returned when the challenge is not found in the store.
	ErrChallengeNotFound = errors.New("challenge not found")

	// ErrChallengeKeyEmpty is returned when the challenge key is empty.
	// This should never happen and could lead to incorrect storage of challenges.
	ErrChallengeKeyEmpty = errors.New("challenge key is empty")

	// ErrChallengeDifficultyInvalid is returned when the challenge difficulty is invalid.
	ErrChallengeDifficultyInvalid = errors.New("challenge difficulty is invalid")

	// ErrChallengeSaltLengthInvalid is returned when the challenge salt length is invalid.
	ErrChallengeSaltLengthInvalid = errors.New("challenge salt length is invalid")
)
