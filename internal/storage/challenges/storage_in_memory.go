package challenges

import (
	"context"

	"go.uber.org/zap"
)

// StorageInMemory is a quote storage in memory.
type StorageInMemory struct {
	logger *zap.Logger
	db     map[string]string
}

// NewStorageInMemory creates a new quote storage in memory.
func NewStorageInMemory(logger *zap.Logger) *StorageInMemory {
	// Logging the call
	logger.Debug("creating a new challenge storage in memory")

	return &StorageInMemory{logger: logger, db: make(map[string]string)}
}

// Add adds a challenge to the store.
func (s *StorageInMemory) Add(ctx context.Context, key, value string) error {
	// Logging the call
	s.logger.Debug("adding challenge to the store", zap.String("key", key), zap.String("value", value))

	// Checking if the context is canceled
	if ctx.Err() != nil {
		return ctx.Err()
	}

	// Adding the challenge to the db
	s.db[key] = value

	// Returning nil as the error
	return nil
}

// Get checks if a challenge is in the store.
func (s *StorageInMemory) Get(ctx context.Context, key string) (bool, error) {
	// Logging the call
	s.logger.Debug("getting challenge from the store", zap.String("key", key))

	// Checking if the context is canceled
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	// Checking if the challenge is in the db
	_, ok := s.db[key]

	// Returning the result and nil as the error
	return ok, nil
}

// Delete deletes a challenge from the store.
func (s *StorageInMemory) Delete(ctx context.Context, key string) error {
	// Logging the call
	s.logger.Debug("deleting challenge from the store", zap.String("key", key))

	// Checking if the context is canceled
	if ctx.Err() != nil {
		return ctx.Err()
	}

	// Deleting the challenge from the db
	delete(s.db, key)

	// Returning nil as the error
	return nil
}
