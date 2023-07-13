package proofer

// Config is the configuration for the proofer middleware.
type Config struct {
	// ChallengeDifficulty is the difficulty of the challenge.
	ChallengeDifficulty int
	// SaltLength is the length of the salt.
	SaltLength int
}
