package ratelimiter

import (
	"net/http"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RateLimiter is a middleware that limits the number of requests a client can make.
type RateLimiter struct {
	logger *zap.Logger
	cfg    *Config
	store  ratelimit.Store
}

// New creates a new rate limiter middleware instance.
func New(logger *zap.Logger, cfg *Config) *RateLimiter {
	return &RateLimiter{
		logger: logger,
		cfg:    cfg,
	}
}

// Use uses the rate limiter middleware.
func (mw *RateLimiter) Use() gin.HandlerFunc {
	// Create a new in memory store with the given options.
	mw.store = ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  mw.parseRate(mw.cfg.Rate),
		Limit: mw.cfg.Limit,
	})

	// Create a new rate limiter middleware instance.
	rateLimiter := ratelimit.RateLimiter(mw.store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      mw.parseKey(mw.cfg.Key),
	})

	// Return the rate limiter middleware.
	return rateLimiter
}

// errorHandler is the function that is called when a request is rejected.
// It returns a 429 status code with a message indicating when the client can retry.
func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(http.StatusTooManyRequests, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}
