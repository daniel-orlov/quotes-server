package mocks

import "context"

// MockChallengeStorage is a mock for challenge storage.
type MockChallengeStorage struct {
	challenges   map[string]string
	storageError error
}

// NewMockChallengeStorage creates a new mock for challenge storage.
func NewMockChallengeStorage(challenges map[string]string, storageError error) *MockChallengeStorage {
	// If no challenges were provided, create a new map
	if challenges == nil {
		challenges = make(map[string]string)
	}

	return &MockChallengeStorage{challenges: challenges, storageError: storageError}
}

// Add adds a challenge to the store.
func (m *MockChallengeStorage) Add(_ context.Context, key, value string) error {
	// check if error was set
	if m.storageError != nil {
		return m.storageError
	}

	// add the challenge to the map
	m.challenges[key] = value

	// return no error
	return nil
}

// Get checks if a challenge is in the store.
func (m *MockChallengeStorage) Get(_ context.Context, key string) (bool, error) {
	// check if error was set
	if m.storageError != nil {
		return false, m.storageError
	}

	// check if the challenge exists in the map
	_, ok := m.challenges[key]

	// return the result
	return ok, nil
}

// Delete deletes a challenge from the store.
func (m *MockChallengeStorage) Delete(_ context.Context, key string) error {
	// check if error was set
	if m.storageError != nil {
		return m.storageError
	}

	// delete the challenge from the map
	delete(m.challenges, key)

	// return no error
	return nil
}
