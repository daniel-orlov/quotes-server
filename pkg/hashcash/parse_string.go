package hashcash

import (
	"crypto/sha1" //nolint:gosec // sha1 is used for hashcash by design
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseStr parses the hashcash string to the hashcash struct, validating it in the process.
// It accepts the hashcash string, expecting it to be in the following format:
// <version>:<difficulty>:<date>:<resource>:<extension>:<salt>:<counter hash>
// It returns a pointer to the hashcash and an error, if any.
func ParseStr(hashcash string) (*Hashcash, error) {
	// Split the hashcash string by ":"
	// The result is an array of strings
	// The first element is the hashcash version
	// The second element is the number of leading zeros required
	// The third element is the date
	// The fourth element is the resource
	// The fifth element is the extension, which is always empty
	// The sixth element is the salt
	// The seventh element is the counter hash
	split := strings.Split(hashcash, ":")

	// Check if the hashcash string has the correct number of elements
	if len(split) != ValidPartsNumber {
		return nil, fmt.Errorf("%w: expected %d, got %d", ErrIncorrectNumberOfParts, ValidPartsNumber, len(split))
	}

	// Convert the hashcash version to int
	version, err := strconv.Atoi(split[0])
	if err != nil {
		return nil, fmt.Errorf("converting hashcash version to int: %w", err)
	}

	// Check if the hashcash version is valid
	if version != Version {
		return nil, fmt.Errorf("%w: expected %d, got %d", ErrInvalidVersion, Version, version)
	}

	// Check if the date format is valid.
	dateFormat, err := ParseDateFormat(split[2])
	if err != nil {
		return nil, fmt.Errorf("parsing hashcash date format: %w", err)
	}

	// Parse the date
	date, err := time.Parse(dateFormat.String(), split[2])
	if err != nil {
		return nil, fmt.Errorf("parsing hashcash date: %w", err)
	}

	// Convert the number of leading zeros required to int
	difficulty, err := strconv.Atoi(split[1])
	if err != nil {
		return nil, fmt.Errorf("converting number of leading zeros required to int: %w", err)
	}

	// Parse the counter_hash back to counter (from hex to int)
	counter, err := strconv.ParseInt(split[6], 16, 32)
	if err != nil {
		return nil, fmt.Errorf("parsing counter hash back to counter: %w", err)
	}

	// Return the hashcash and nil error
	return &Hashcash{
		hash:       sha1.New(), //nolint:gosec // sha1 is used for hashcash by design
		version:    version,
		difficulty: difficulty,
		date:       date,
		dateFormat: dateFormat,
		resource:   split[3],
		salt:       split[5],
		counter:    int(counter),
	}, nil
}
