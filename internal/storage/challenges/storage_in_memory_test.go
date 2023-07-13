package challenges_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/internal/storage/challenges"
)

func TestStorageInMemory_Add(t *testing.T) {
	t.Run("add challenge to the store", func(t *testing.T) {
		// create a new storage
		store := challenges.NewStorageInMemory(zap.NewNop())

		// add a challenge to the store
		err := store.Add(context.Background(), "key", "value")

		// check that there is no error
		assert.NoError(t, err, "there should be no error")

		// check that the challenge is in the store
		ok, err := store.Get(context.Background(), "key")

		// check that there is no error
		assert.NoError(t, err, "there should be no error")

		// check that the challenge is in the store
		assert.True(t, ok, "the challenge should be in the store")
	})

	t.Run("add challenge to the store with a canceled context", func(t *testing.T) {
		// create a new storage
		store := challenges.NewStorageInMemory(zap.NewNop())

		// create a canceled context
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		// add a challenge to the store
		err := store.Add(ctx, "key", "value")

		// check that there is an error
		assert.Error(t, err, "there should be an error")

		// check that the challenge is not in the store
		ok, err := store.Get(context.Background(), "key")

		// check that there is no error
		assert.NoError(t, err, "there should be no error")

		// check that the challenge is not in the store
		assert.False(t, ok, "the challenge should not be in the store")
	})
}

func TestStorageInMemory_Delete(t *testing.T) {
	t.Run("delete challenge from the store", func(t *testing.T) {
		// create a new storage
		store := challenges.NewStorageInMemory(zap.NewNop())

		// add a challenge to the store
		err := store.Add(context.Background(), "key", "value")

		// check that there is no error
		assert.NoError(t, err, "there should be no error")

		// delete the challenge from the store
		err = store.Delete(context.Background(), "key")

		// check that there is no error
		assert.NoError(t, err, "there should be no error")

		// check that the challenge is not in the store
		ok, err := store.Get(context.Background(), "key")

		// check that there is no error
		assert.NoError(t, err, "there should be no error")

		// check that the challenge is not in the store
		assert.False(t, ok, "the challenge should not be in the store")
	})

	t.Run("delete challenge from the store with a canceled context", func(t *testing.T) {
		// create a new storage
		store := challenges.NewStorageInMemory(zap.NewNop())

		// add a challenge to the store
		err := store.Add(context.Background(), "key", "value")

		// check that there is no error
		assert.NoError(t, err, "there should be no error")

		// create a canceled context
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		// delete the challenge from the store
		err = store.Delete(ctx, "key")

		// check that there is an error
		assert.Error(t, err, "there should be an error")

		// check that the challenge is in the store
		ok, err := store.Get(context.Background(), "key")

		// check that there is no error
		assert.NoError(t, err, "there should be no error")

		// check that the challenge is in the store
		assert.True(t, ok, "the challenge should be in the store")
	})
}

func TestStorageInMemory_Get(t *testing.T) {
	t.Run("get challenge from the store", func(t *testing.T) {
		// create a new storage
		store := challenges.NewStorageInMemory(zap.NewNop())

		// add a challenge to the store
		err := store.Add(context.Background(), "key", "value")

		// check that there is no error
		assert.NoError(t, err, "there should be no error")

		// get the challenge from the store
		ok, err := store.Get(context.Background(), "key")

		// check that there is no error
		assert.NoError(t, err, "there should be no error")

		// check that the challenge is in the store
		assert.True(t, ok, "the challenge should be in the store")
	})

	t.Run("get challenge from the store with a non-existent key", func(t *testing.T) {
		// create a new storage
		store := challenges.NewStorageInMemory(zap.NewNop())

		// get the challenge from the store
		ok, err := store.Get(context.Background(), "key")

		// check that there is no error
		assert.NoError(t, err, "there should be no error")

		// check that the challenge is not in the store
		assert.False(t, ok, "the challenge should not be in the store")
	})

	t.Run("get challenge from the store with a canceled context", func(t *testing.T) {
		// create a new storage
		store := challenges.NewStorageInMemory(zap.NewNop())

		// add a challenge to the store
		err := store.Add(context.Background(), "key", "value")

		// check that there is no error
		assert.NoError(t, err, "there should be no error")

		// create a canceled context
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		// get the challenge from the store
		ok, err := store.Get(ctx, "key")

		// check that there is an error
		assert.Error(t, err, "there should be an error")

		// check that the challenge is in the store
		assert.False(t, ok, "the challenge should not be in the store")
	})
}
