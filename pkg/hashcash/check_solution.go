package hashcash

import "fmt"

// Check checks if the hashcash is solved and not expired.
func (h *Hashcash) Check() (bool, error) {
	// Check if the hashcash has expired
	expired, err := h.HasExpired()
	if err != nil {
		return false, fmt.Errorf("checking if hashcash has expired: %w", err)
	}

	if expired {
		return false, ErrExpiredHashcash
	}

	// Check if the hashcash is solved
	solved := h.IsSolved()

	// If the hashcash is not solved, return error
	if !solved {
		return false, ErrIncorrectSolution
	}

	// Solution is correct
	return true, nil
}

// CheckSolution parses solution string and checks if the hashcash is solved and not expired.
func CheckSolution(hashcashStr string) (bool, error) {
	// Parse the hashcash string
	hc, err := ParseStr(hashcashStr)
	if err != nil {
		return false, fmt.Errorf("parsing hashcash string: %w", err)
	}

	// Check if the hashcash is solved and not expired
	return hc.Check()
}
