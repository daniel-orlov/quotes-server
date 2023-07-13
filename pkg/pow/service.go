package pow

import (
	"context"

	"go.uber.org/zap"
)

// ChallengeStore is an interface for storing which challenges are currently active.
type ChallengeStore interface {
	// Add adds a challenge to the store.
	Add(ctx context.Context, key, value string) error
	// Get checks if a challenge is in the store.
	Get(ctx context.Context, key string) (bool, error)
	// Delete deletes a challenge from the store.
	Delete(ctx context.Context, key string) error
}

// Service is a PoW service.
type Service struct {
	logger *zap.Logger
	// ChallengeStore is a store for challenges.
	Store ChallengeStore
}

// NewService creates a new PoW service.
func NewService(logger *zap.Logger, store ChallengeStore) *Service {
	return &Service{logger: logger, Store: store}
}
