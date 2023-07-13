package pow

import (
	"context"
	"fmt"

	"github.com/daniel-orlov/quotes-server/pkg/hashcash"
)

// CheckSolution checks if the solution with the given challenge key exists in the store and if it is correct.
func (s *Service) CheckSolution(ctx context.Context, solution string, key Key) (bool, error) {
	// If the key is nil, return error
	if key == nil {
		return false, ErrChallengeKeyEmpty
	}

	// Stringify the key
	stringKey := key.String()

	// If the key is empty, return error
	if stringKey == ":" {
		return false, ErrChallengeKeyEmpty
	}

	// Check if the solution exists in the store
	exists, err := s.Store.Get(ctx, stringKey)
	if err != nil {
		return false, fmt.Errorf("getting challenge from store: %w", err)
	}

	// If the solution does not exist in the store, return error
	if !exists {
		return false, ErrChallengeNotFound
	}

	// Check if the solution is correct
	correct, err := hashcash.CheckSolution(solution)
	if err != nil {
		return false, fmt.Errorf("checking solution: %w", err)
	}

	// if the solution is correct, delete it from the store
	if correct {
		err = s.Store.Delete(ctx, stringKey)
		if err != nil {
			return false, fmt.Errorf("deleting challenge from store: %w", err)
		}
	}

	// Return the result
	return correct, nil
}
