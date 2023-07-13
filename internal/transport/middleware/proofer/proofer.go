// Package proofer provides a middleware that implements a proof of work challenge.
// The middleware is used to prevent DoS attacks.
// It requires the client to solve a proof of work challenge before the request is processed.
package proofer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/pkg/pow"
)

// PoWService is a port to the PoW service.
type PoWService interface {
	NewChallenge(ctx context.Context, challengeKey pow.Key, difficulty, saltLength int) (string, error)
	CheckSolution(ctx context.Context, solution string, challengeKey pow.Key) (bool, error)
}

// Proofer is a middleware that checks Proof-of-Work in request and thus prevents DoS-attacks.
type Proofer struct {
	logger *zap.Logger
	cfg    *Config
	svc    PoWService
}

// New creates new Proofer middleware.
func New(logger *zap.Logger, cfg *Config, svc PoWService) *Proofer {
	// Logging the call
	logger.Debug("creating a new proofer middleware")

	return &Proofer{logger: logger, cfg: cfg, svc: svc}
}

const (
	// ChallengeHeader is the name of the header that contains the challenge solution, like a hashcash challenge.
	ChallengeHeader = "X-Hashcash"
)

// Use uses the proofer middleware.
func (mw *Proofer) Use() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the solution from the request
		solution := c.GetHeader(ChallengeHeader)

		// if hashcash is not present, return challenge
		if solution == "" {
			// Try to get a new challenge
			if err := mw.handleNewChallengeRequest(c); err != nil {
				// Return error
				mw.handleError(c, http.StatusInternalServerError, "failed to get new challenge", err)
				// Abort request
				return
			}

			// Abort request, return challenge
			mw.abortRequest(c, http.StatusPreconditionRequired, "proof-of-work requirements not met")
			return
		}

		// Try to check the solution
		if err := mw.handleSolutionCheck(c, solution); err != nil {
			// Return error
			mw.handleError(c, http.StatusInternalServerError, "failed to check solution", err)
			return
		}

		// Continue processing the request
		c.Next()
	}
}

// handleNewChallengeRequest handles a request for a new challenge.
func (mw *Proofer) handleNewChallengeRequest(c *gin.Context) error {
	// Get a new challenge from the service
	challenge, err := mw.svc.NewChallenge(
		c.Request.Context(),
		pow.NewChallengeKey(c.ClientIP(), fmt.Sprintf("%s:%s", c.Request.Method, c.Request.URL.Path)),
		mw.cfg.ChallengeDifficulty,
		mw.cfg.SaltLength,
	)
	// Handle error
	if err != nil {
		mw.logger.Error("failed to get new challenge", zap.Error(err))
		return err
	}

	// Set the challenge in the response header
	c.Header(ChallengeHeader, challenge)

	// Challenge is set, return nil
	return nil
}

// handleSolutionCheck handles a request for a solution check.
func (mw *Proofer) handleSolutionCheck(c *gin.Context, solution string) error {
	// Check the solution with the service
	solved, err := mw.svc.CheckSolution(
		c.Request.Context(),
		solution,
		pow.NewChallengeKey(c.ClientIP(), fmt.Sprintf("%s:%s", c.Request.Method, c.Request.URL.Path)),
	)
	// Handle error
	if err != nil {
		mw.logger.Error("failed to check solution", zap.Error(err))
		return err
	}

	// If solution is not valid, return error and a new challenge, abort request
	if !solved {
		// Try to get a new challenge
		if err = mw.handleNewChallengeRequest(c); err != nil {
			// Return error
			mw.handleError(c, http.StatusInternalServerError, "failed to get new challenge", err)
			return err
		}

		// Abort request, return challenge
		mw.abortRequest(c, http.StatusPreconditionRequired, "proof-of-work requirements not met: solution is invalid")
	}

	// Solution is valid, return nil
	return nil
}

// handleError handles an error.
func (mw *Proofer) handleError(c *gin.Context, statusCode int, errorMessage string, err error) {
	// Log the actual error
	mw.logger.Error(errorMessage, zap.Error(err))

	// Abort with a generic error message
	c.AbortWithStatusJSON(statusCode, gin.H{
		"error": errorMessage,
	})
}

// abortRequest aborts the request with a generic error message.
// Don't use this method to abort the request with an actual error message.
func (mw *Proofer) abortRequest(c *gin.Context, statusCode int, errorMessage string) {
	c.AbortWithStatusJSON(statusCode, gin.H{
		"error": errorMessage,
	})
}
