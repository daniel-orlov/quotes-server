package pow

import "fmt"

// Key is a key for uniquely identifying a challenge.
type Key interface {
	String() string
	ClientID() string
	ResourceID() string
}

// ChallengeKey is a key for a challenge.
type ChallengeKey struct {
	// clientID is the client ID.
	// Could be a UUID, some hash or client's IP address.
	clientID string
	// resourceID is the resource ID.
	// Could be a UUID, some hash or resource's URL.
	resourceID string
}

// NewChallengeKey returns a new challenge key.
func NewChallengeKey(clientID string, resourceID string) *ChallengeKey {
	return &ChallengeKey{clientID: clientID, resourceID: resourceID}
}

// String returns a string representation of the challenge key.
func (c *ChallengeKey) String() string {
	// Check if the challenge key is nil to avoid panics
	if c == nil {
		return ""
	}

	// Return the challenge key as a string
	return fmt.Sprintf("%s:%s", c.ClientID(), c.ResourceID())
}

// ClientID returns the client ID.
func (c *ChallengeKey) ClientID() string {
	// Check if the challenge key is nil to avoid panics
	if c == nil {
		return ""
	}

	// Return the client ID
	return c.clientID
}

// ResourceID returns the resource ID.
func (c *ChallengeKey) ResourceID() string {
	// Check if the challenge key is nil to avoid panics
	if c == nil {
		return ""
	}

	// Return the resource ID
	return c.resourceID
}
