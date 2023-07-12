package hashcash_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daniel-orlov/quotes-server/pkg/hashcash"
)

func TestParseStr(t *testing.T) {
	// Format: <version>:<difficulty>:<date>:<resource>:<extension>:<salt>:<counter hash>

	t.Run("invalid string - incorrect number of elements", func(t *testing.T) {
		// Parse invalid hashcash string
		_, err := hashcash.ParseStr("1:20:060102150405::salt")

		// Check if the error is returned
		assert.ErrorIs(t, err, hashcash.ErrIncorrectNumberOfParts, "incorrect number of parts: expected 7, got 5")
	})

	t.Run("invalid string - incorrect hashcash version", func(t *testing.T) {
		// Parse invalid hashcash string
		_, err := hashcash.ParseStr("2:20:060102150405:some-resource::salt:counter_hash")

		// Check if the error is returned
		assert.ErrorIs(t, err, hashcash.ErrInvalidVersion, "hashcash version is invalid: expected 1, got 2")
	})

	t.Run("invalid string - invalid hashcash version", func(t *testing.T) {
		// Parse invalid hashcash string
		_, err := hashcash.ParseStr("NaN:20:060102150405:resource::salt:counter_hash")

		// Check if the error is returned
		assert.Error(t, err, "hashcash version is invalid, expected 1, got NaN")
	})

	t.Run("invalid string - incorrect date format", func(t *testing.T) {
		// Parse invalid hashcash string
		_, err := hashcash.ParseStr("1:20:12345:resource::salt:counter_hash")

		// Check if the error is returned
		assert.Error(t, err, "parsing date should fail")
	})

	t.Run("invalid string - incorrect number of leading zeros required", func(t *testing.T) {
		// Parse invalid hashcash string
		_, err := hashcash.ParseStr("1:NaN:06:resource::salt:counter_hash")

		// Check if the error is returned
		assert.Error(t, err, "converting number of leading zeros required to int: strconv.Atoi: parsing \"NaN\": invalid syntax")
	})

	t.Run("invalid string - incorrect counter hash", func(t *testing.T) {
		// Parse invalid hashcash string
		_, err := hashcash.ParseStr("1:20:0212:resource::salt:NaN")

		// Check if the error is returned
		assert.Error(t, err, "parsing counter hash back to counter: strconv.ParseInt: parsing \"NaN\": invalid syntax")
	})

	t.Run("valid string", func(t *testing.T) {
		// Parse valid hashcash string
		validHCString := "1:20:041010:resource::salt:23a"

		hc, err := hashcash.ParseStr(validHCString)

		// Check that no error is returned
		assert.NoError(t, err)

		// Check if the hashcash is correct
		assert.Equal(t, validHCString, hc.String())
	})
}
