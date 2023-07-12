package hashcash_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daniel-orlov/quotes-server/pkg/hashcash"
)

func TestCheckSolution(t *testing.T) {
	// <version>:<difficulty>:<date>:<resource>:<extension>:<salt>:<counter hash>

	t.Run("parsing failed", func(t *testing.T) {
		// Invalid hashcash
		invalidHC := "invalid hashcash"

		// Checking solution should fail
		passed, err := hashcash.CheckSolution(invalidHC)

		// Error should not be nil
		assert.Error(t, err, "parsing failed")

		// Solution should not pass the check
		assert.False(t, passed, "solution should not pass")
	})

	t.Run("checking expiration failed", func(t *testing.T) {
		// Hashcash with invalid date
		expiredHC := "1:20:000000000000:some-resource::salt:23a"

		// Checking solution should fail
		passed, err := hashcash.CheckSolution(expiredHC)

		// Error should not be nil
		assert.Error(t, err, "checking expiration failed")

		// Solution should not pass the check
		assert.False(t, passed, "solution should not pass")
	})

	t.Run("expired hashcash", func(t *testing.T) {
		// Hashcash with expired date
		expiredHC := "1:20:060102150405:some-resource::salt:23a"

		// Checking solution should fail
		passed, err := hashcash.CheckSolution(expiredHC)

		// Error should not be nil
		assert.Error(t, err, "checking expiration failed")

		// Solution should not pass the check
		assert.False(t, passed, "solution should not pass")
	})

	t.Run("hashcash is not solved", func(t *testing.T) {
		// Valid hashcash, but not solved
		unsolvedHC := "1:20:23:some-resource::salt:23a"

		// Checking solution should fail
		passed, err := hashcash.CheckSolution(unsolvedHC)

		// Error should be nil
		assert.ErrorIs(t, err, hashcash.ErrIncorrectSolution, "incorrect solution")

		// Solution should not pass the check
		assert.False(t, passed, "solution should not pass")
	})

	t.Run("hashcash is solved", func(t *testing.T) {
		// Valid hashcash, solved
		validAndSolvedHC := "1:20:23:some-resource::Kl7oUEQg:4c73d"

		// Checking solution should pass
		passed, err := hashcash.CheckSolution(validAndSolvedHC)

		// Error should be nil
		assert.NoError(t, err, "checking solution failed")

		// Solution should pass the check
		assert.True(t, passed, "solution should pass")
	})
}
