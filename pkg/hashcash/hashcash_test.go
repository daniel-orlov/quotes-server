package hashcash_test

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/daniel-orlov/quotes-server/pkg/hashcash"
)

func Test_New(t *testing.T) {
	t.Run("difficulty is zero, should return error", func(t *testing.T) {
		// Create a new hashcash
		_, err := hashcash.New(0, 8, hashcash.DateFormatYYMMDD, "resource")

		// Check if the error is nil
		assert.Error(t, err, "creating hashcash should return an error")
	})

	t.Run("salt length is zero, should return error", func(t *testing.T) {
		// Create a new hashcash
		_, err := hashcash.New(20, 0, hashcash.DateFormatYYMMDD, "resource")

		// Check if the error is nil
		assert.Error(t, err, "creating hashcash should return an error")
	})

	t.Run("date format is empty, should return error", func(t *testing.T) {
		// Create a new hashcash
		_, err := hashcash.New(20, 8, "", "resource")

		// Check if the error is nil
		assert.Error(t, err, "creating hashcash should return an error")
	})

	t.Run("resource is empty, should not return error", func(t *testing.T) {
		// Create a new hashcash
		_, err := hashcash.New(20, 8, hashcash.DateFormatYYMMDD, "")

		// Check if the error is nil
		assert.NoError(t, err, "creating hashcash should not return an error")
	})

	t.Run("all parameters are valid, should not return error", func(t *testing.T) {
		// Create a new hashcash
		hc, err := hashcash.New(20, 8, hashcash.DateFormatYYMMDD, "resource")

		// Check if the error is nil
		assert.NoError(t, err, "creating hashcash should not return an error")

		// Check if the hashcash is not nil
		assert.NotNil(t, hc, "hashcash should not be nil")
	})
}

func TestHashcash_String(t *testing.T) {
	t.Run("hashcash is nil, should return empty string", func(t *testing.T) {
		// Create a new nil hashcash
		var hc *hashcash.Hashcash

		// Get the hashcash string
		stringHC := hc.String()

		// Check if the hashcash string is empty
		assert.Equal(t, "", stringHC, "hashcash string should be empty")
	})

	t.Run("hashcash is empty, should return string with zero values", func(t *testing.T) {
		// Create a new empty hashcash
		hc := hashcash.Hashcash{}

		// Get the hashcash string
		stringHC := hc.String()

		// Check if the hashcash string is empty
		assert.Equal(t, "0:0:::::0", stringHC, "hashcash string should be empty")
	})

	t.Run("hashcash is not empty, should return correct string", func(t *testing.T) {
		// Prepare the expected hashcash string
		now := time.Now()
		saltLen := 8
		difficulty := 20
		difficultyString := "20"
		versionString := "1"

		// Create a new hashcash
		hc, err := hashcash.New(difficulty, saltLen, hashcash.DateFormatYYMMDD, "resource")
		assert.NoError(t, err, "creating hashcash should not return an error")

		// Get the hashcash string
		stringHC := hc.String()

		// Since we shouldn't be testing implementation of counter hash and salt here,
		// we can split the hashcash string and compare only the parts that we expect to be the same

		// Split the hashcash string
		splitStringHC := strings.Split(stringHC, ":")

		// Check if the hashcash string has the correct number of parts
		assert.Equal(t, hashcash.ValidPartsNumber, len(splitStringHC), "hashcash string should have the correct number of parts")

		// Check if the hashcash string has the correct version
		assert.Equal(t, versionString, splitStringHC[0], "hashcash string should have the correct version")

		// Check if the hashcash string has the correct difficulty
		assert.Equal(t, difficultyString, splitStringHC[1], "hashcash string should have the correct difficulty")

		// Check if the hashcash string has the correct date
		assert.Equal(t, now.Format(hashcash.DateFormatYYMMDD.String()), splitStringHC[2], "hashcash string should have the correct date")

		// Check if the hashcash string has the correct resource
		assert.Equal(t, "resource", splitStringHC[3], "hashcash string should have the correct resource")

		// Check that the length of the salt is correct
		assert.Equal(t, saltLen, len(splitStringHC[5]), "hashcash string should have the correct salt length")
	})
}

func TestHashcash_HasExpired(t *testing.T) {
	t.Run("hashcash is nil, should return error", func(t *testing.T) {
		// Create a new nil hashcash
		var hc *hashcash.Hashcash

		// Check if the hashcash has expired
		_, err := hc.HasExpired()

		// Check if the error is not nil
		assert.Error(t, err, "hashcash should not be nil")
	})

	t.Run("hashcash is empty, should return true and no error", func(t *testing.T) {
		// Create a new empty hashcash
		hc := hashcash.Hashcash{}

		// Check if the hashcash has expired
		hasExpired, err := hc.HasExpired()

		// Check if the error is nil
		assert.NoError(t, err)

		// Check if the hashcash has expired
		assert.True(t, hasExpired, "hashcash should have expired")
	})

	t.Run("hashcash is not empty but date is in the future, should return false", func(t *testing.T) {
		// Create a new hashcash
		hc, err := hashcash.ParseStr("1:20:33:resource::salt:23a")

		// Check if the error is nil
		assert.NoError(t, err, "parsing hashcash should not return an error")

		// Check if the hashcash has expired
		hasExpired, err := hc.HasExpired()

		// Should return error
		assert.ErrorIs(t, err, hashcash.ErrAttemptToUseFutureHashcash)

		// Check if the hashcash has expired
		assert.False(t, hasExpired, "hashcash should be invalid")
	})

	t.Run("hashcash is not empty and not expired (date format is DateFormatYY), should return false", func(t *testing.T) {
		// Create a new hashcash
		hc, err := hashcash.New(20, 8, hashcash.DateFormatYY, "resource")
		assert.NoError(t, err, "creating hashcash should not return an error")

		// Check if the hashcash has expired
		hasExpired, err := hc.HasExpired()

		// Check if the error is nil
		assert.NoError(t, err)

		// Check if the hashcash has expired
		assert.False(t, hasExpired, "hashcash should not have expired")
	})

	t.Run("hashcash is not empty and not expired (date format is DateFormatYYMM), should return false", func(t *testing.T) {
		// Create a new hashcash
		hc, err := hashcash.New(20, 8, hashcash.DateFormatYYMM, "resource")
		assert.NoError(t, err, "creating hashcash should not return an error")

		// Check if the hashcash has expired
		hasExpired, err := hc.HasExpired()

		// Check if the error is nil
		assert.NoError(t, err)

		// Check if the hashcash has expired
		assert.False(t, hasExpired, "hashcash should not have expired")
	})

	t.Run("hashcash is not empty and not expired (date format is DateFormatYYMMDD), should return false", func(t *testing.T) {
		// Create a new hashcash
		hc, err := hashcash.New(20, 8, hashcash.DateFormatYYMMDD, "resource")
		assert.NoError(t, err, "creating hashcash should not return an error")

		// Check if the hashcash has expired
		hasExpired, err := hc.HasExpired()

		// Check if the error is nil
		assert.NoError(t, err)

		// Check if the hashcash has expired
		assert.False(t, hasExpired, "hashcash should not have expired")
	})

	t.Run("hashcash is not empty and not expired (date format is DateFormatYYMMDDhhmm), should return false", func(t *testing.T) {
		// Create a new hashcash
		hc, err := hashcash.New(20, 8, hashcash.DateFormatYYMMDDhhmm, "resource")
		assert.NoError(t, err, "creating hashcash should not return an error")

		// Check if the hashcash has expired
		hasExpired, err := hc.HasExpired()

		// Check if the error is nil
		assert.NoError(t, err)

		// Check if the hashcash has expired
		assert.False(t, hasExpired, "hashcash should not have expired")
	})

	t.Run("hashcash is not empty and not expired (date format is DateFormatYYMMDDhhmmss), should return false", func(t *testing.T) {
		// Create a new hashcash
		hc, err := hashcash.New(20, 8, hashcash.DateFormatYYMMDDhhmmss, "resource")
		assert.NoError(t, err, "creating hashcash should not return an error")

		// Check if the hashcash has expired
		hasExpired, err := hc.HasExpired()

		// Check if the error is nil
		assert.NoError(t, err)

		// Check if the hashcash has expired
		assert.False(t, hasExpired, "hashcash should not have expired")
	})
}

func TestHashcash_IsSolved(t *testing.T) {
	t.Run("hashcash is nil, should return false", func(t *testing.T) {
		// Create a new nil hashcash
		var hc *hashcash.Hashcash

		// Check if the hashcash is solved
		isSolved := hc.IsSolved()

		// Assert that the hashcash is not solved
		assert.False(t, isSolved, "hashcash should not be solved")
	})

	t.Run("hashcash is empty, should return false", func(t *testing.T) {
		// Create a new empty hashcash
		hc := hashcash.Hashcash{}

		// Check if the hashcash is solved
		isSolved := hc.IsSolved()

		// Assert that the hashcash is not solved
		assert.False(t, isSolved, "hashcash should not be solved")
	})

	t.Run("hashcash is not empty and not solved, should return false", func(t *testing.T) {
		// Create a new hashcash
		hc, err := hashcash.New(20, 8, hashcash.DateFormatYYMMDD, "resource")
		assert.NoError(t, err, "creating hashcash should not return an error")

		// Check if the hashcash is solved
		isSolved := hc.IsSolved()

		// Assert that the hashcash is not solved
		assert.False(t, isSolved, "hashcash should not be solved")
	})

	t.Run("hashcash is not empty and solved, should return true", func(t *testing.T) {
		// Create a new hashcash
		hc, err := hashcash.New(10, 8, hashcash.DateFormatYYMMDD, "resource")
		assert.NoError(t, err, "creating hashcash should not return an error")

		// Solve the hashcash
		_, err = hc.Solve()
		assert.NoError(t, err, "solving hashcash should not return an error")

		// Check if the hashcash has expired
		isSolved := hc.IsSolved()

		// Assert that the hashcash is not solved
		assert.True(t, isSolved, "hashcash should be solved")
	})
}

func TestHashcash_Solve(t *testing.T) {
	t.Run("hashcash is nil, should return error", func(t *testing.T) {
		// Create a new nil hashcash
		var hc *hashcash.Hashcash

		// Solve the hashcash
		solution, err := hc.Solve()

		// Assert that the error is not nil
		assert.Error(t, err, "hashcash should not be nil")

		// Assert that the solution is empty
		assert.Equal(t, "", solution, "hashcash should be empty")
	})

	t.Run("hashcash is empty, should return error", func(t *testing.T) {
		// Create a new empty hashcash
		hc := hashcash.Hashcash{}

		// Solve the hashcash
		solution, err := hc.Solve()

		// Assert that the error is not nil
		assert.Error(t, err, "hashcash should not be empty")

		// Assert that the solution is empty
		assert.Equal(t, "", solution, "hashcash should be empty")
	})

	t.Run("hashcash is not empty and not solved, should return no error", func(t *testing.T) {
		// Create a new hashcash
		hc, err := hashcash.New(20, 8, hashcash.DateFormatYYMMDD, "resource")
		assert.NoError(t, err, "creating hashcash should not return an error")

		// Solve the hashcash
		solution, err := hc.Solve()

		// Assert that the error is nil
		assert.NoError(t, err, "hashcash should not be nil")

		// Assert that the solution is not empty
		assert.NotEqual(t, "", solution, "hashcash should not be empty")
	})
}
