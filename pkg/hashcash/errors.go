package hashcash

import "errors"

var (
	// ErrExpiredHashcash is returned when hashcash is expired.
	ErrExpiredHashcash = errors.New("hashcash is expired")

	// ErrNilHashcash is returned when hashcash is nil.
	ErrNilHashcash = errors.New("hashcash is nil")

	// ErrEmptyHashcash is returned when hashcash is empty.
	ErrEmptyHashcash = errors.New("hashcash is empty")

	// ErrIncorrectNumberOfParts is returned when the number of parts contained in the hashcash string is incorrect.
	ErrIncorrectNumberOfParts = errors.New("incorrect number of parts")

	// ErrInvalidVersion is returned when the hashcash version is invalid.
	ErrInvalidVersion = errors.New("hashcash version is invalid")

	// ErrInvalidSaltLength is returned when the salt length is invalid.
	ErrInvalidSaltLength = errors.New("invalid salt length")

	// ErrInvalidDifficulty is returned when the difficulty is invalid.
	ErrInvalidDifficulty = errors.New("invalid difficulty")

	// ErrInvalidDateFormat is returned when the date format is invalid.
	ErrInvalidDateFormat = errors.New("invalid date format")

	// ErrIncorrectSolution is returned when the hashcash solution is incorrect.
	ErrIncorrectSolution = errors.New("incorrect solution")

	// ErrAttemptToUseFutureHashcash is returned when the hashcash date is in the future.
	ErrAttemptToUseFutureHashcash = errors.New("attempt to use future hashcash")
)
