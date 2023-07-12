package hashcash

import (
	"crypto/rand"
	"crypto/sha1" //nolint:gosec // sha1 is used for hashcash by design
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"hash"
	"strconv"
	"time"
)

// Hashcash is a representation of a hashcash version 1.
// Read more:
// - https://en.wikipedia.org/wiki/Hashcash
// - https://www.hashcash.org/
type Hashcash struct {
	// hash is the hashcash hash function.
	// Currently, it is always sha1.
	// It is added to the struct to avoid creating a new hash function every time.
	hash hash.Hash

	// Version is the hashcash version.
	// It is always 1, as of this writing.
	version int

	// Difficulty is the number of leading zeros required.
	difficulty int

	// Date is the date of the hashcash creation.
	date time.Time

	// DateFormat is the date format used in hashcash.
	dateFormat DateFormat

	// Resource is the resource to which the hashcash is tied.
	// It could be an email address, a domain name, client IP address, or anything else.
	resource string

	// Extension was deprecated in version 1, so it is always empty.
	// It made no sense to include it in the struct.

	// Salt is a random string of characters.
	// It is used to prevent hashcash collisions.
	salt string

	// Counter is a nonce.
	// It is incremented until the hashcash is valid, i.e. has the required number of leading zeros.
	counter int
}

const (
	// Version is the hashcash version.
	// It is always 1, as of this writing.
	Version = 1

	// ValidPartsNumber is the number of parts in a valid hashcash string.
	ValidPartsNumber = 7
)

// New creates a new hashcash.
// It accepts the number of leading zeros and the resource.
// It returns a pointer to the hashcash and an error, if any.
func New(difficulty, saltLen int, dateFormat DateFormat, resource string) (*Hashcash, error) {
	// Check if the difficulty is valid: it should be greater than 0.
	if difficulty <= 0 {
		return nil, ErrInvalidDifficulty
	}

	// Check if the date format is valid.
	if !dateFormat.IsValid() {
		return nil, ErrInvalidDateFormat
	}

	// Create salt of the given length
	// Handle the error, if any
	salt, err := newSalt(saltLen)
	if err != nil {
		return nil, fmt.Errorf("creating salt: %w", err)
	}

	// Return the hashcash and nil error
	return &Hashcash{
		hash:       sha1.New(), //nolint:gosec // sha1 is used for hashcash by design
		version:    Version,
		difficulty: difficulty,
		date:       time.Now(),
		dateFormat: dateFormat,
		salt:       salt,
		resource:   resource,
		// counter is 0 by default
	}, nil
}

// HasExpired checks if the hashcash has expired.
func (h *Hashcash) HasExpired() (bool, error) {
	if h == nil {
		return false, ErrNilHashcash
	}

	// Get the duration since the hashcash date
	duration := time.Since(h.date)

	// Check that the duration is not positive
	if duration < 0 {
		return false, ErrAttemptToUseFutureHashcash
	}

	// Assume the hashcash has expired
	expired := true

	// Switch the hashcash date format
	switch h.dateFormat {
	case DateFormatYY:
		// Can not expire
		expired = false
	case DateFormatYYMM:
		expired = duration > MaxDurationYYMM
	case DateFormatYYMMDD:
		expired = duration > MaxDurationYYMMDD
	case DateFormatYYMMDDhhmm:
		expired = duration > MaxDurationYYMMDDhhmm
	case DateFormatYYMMDDhhmmss:
		expired = duration > MaxDurationYYMMDDhhmmss
	}

	// Return the expiration status and nil error
	return expired, nil
}

// IsSolved checks if the hashcash is solved.
func (h *Hashcash) IsSolved() bool {
	// Return false if the hashcash is nil. This is to avoid panics.
	if h == nil || h.hash == nil {
		return false
	}

	// Reset the hash function
	h.hash.Reset()

	// Write the hashcash string to the hash
	h.hash.Write([]byte(h.String()))

	// Get the hash sum
	sum := h.hash.Sum(nil)

	// Cast the hash sum to uint64
	sumUint64 := binary.BigEndian.Uint64(sum)

	// Convert the hash sum to binary string
	// This effectively removes the leading zeros
	sumBits := strconv.FormatUint(sumUint64, 2)

	// Get the number of leading zeros
	leadingZeroes := 64 - len(sumBits)

	// Check the number of leading zeros
	return leadingZeroes >= h.difficulty
}

// Solve solves the hashcash using brute force.
func (h *Hashcash) Solve() (string, error) {
	if h == nil {
		return "", ErrNilHashcash
	}

	if h.hash == nil || h.date.IsZero() || h.resource == "" || h.salt == "" {
		return "", ErrEmptyHashcash
	}

	// Increment the hashcash counter until it is valid
	for !h.IsSolved() {
		// Increment the counter
		h.counter++
	}

	// Return the hashcash and nil error
	return h.String(), nil
}

// String returns the hashcash string.
func (h *Hashcash) String() string {
	// Return an empty string if the hashcash is nil. This is to avoid panics.
	if h == nil {
		return ""
	}

	return fmt.Sprintf(
		"%d:%d:%s:%s::%s:%x",
		h.version,
		h.difficulty,
		h.date.Format(h.dateFormat.String()), // Convert date to string of the given dateFormat
		h.resource,
		h.salt,
		h.counter,
	)
}

// newSalt creates a salt of the given length.
func newSalt(saltLen int) (string, error) {
	// Check if the salt length is valid
	if saltLen <= 0 {
		return "", ErrInvalidSaltLength
	}

	// Create a byte slice of the given length
	buf := make([]byte, saltLen)

	// Read random bytes into the byte slice
	_, err := rand.Read(buf)
	// Handle the error
	if err != nil {
		return "", fmt.Errorf("reading random bytes: %w", err)
	}

	// Encode the byte slice to base64
	salt := base64.StdEncoding.EncodeToString(buf)

	// Return the salt
	return salt[:saltLen], nil
}
