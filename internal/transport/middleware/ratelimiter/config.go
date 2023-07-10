package ratelimiter

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Config is the configuration for the rate limiter middleware.
type Config struct {
	// Rate is the rate at which requests are allowed. Could be specified as Second, Minute, Hour or Day.
	Rate Rate
	// Limit is the maximum number of requests that can be made in the given Rate.
	Limit uint
	// Key is the way to identify the client. By default, it uses the client's IP address.
	Key Key
}

// Key is the type used to identify keys in the store.
type Key string

const (
	// ClientIP is the key used in the store to identify the client.
	ClientIP Key = "client_ip"
)

// Rate is the type used to specify the rate at which requests are allowed.
type Rate string

const (
	// Second is the rate at which requests are allowed per second.
	Second Rate = "second"
	// Minute is the rate at which requests are allowed per minute.
	Minute Rate = "minute"
)

// parseRate parses the rate string and returns the time.Duration equivalent.
func (mw *RateLimiter) parseRate(rate Rate) time.Duration {
	// to lower case
	rate = Rate(strings.ToLower(string(rate)))

	// switch on rate and return the time.Duration equivalent
	switch rate {
	case Second:
		return time.Second
	case Minute:
		return time.Minute
	default:
		// log warning if unknown rate is used
		mw.logger.Warn("unknown rate, using default",
			zap.String("rate", string(rate)),
			zap.String("default", string(Second)),
		)

		return time.Second
	}
}

func (mw *RateLimiter) parseKey(key Key) func(c *gin.Context) string {
	// to lower case
	key = Key(strings.ToLower(string(key)))

	switch key {
	case ClientIP:
		return func(c *gin.Context) string {
			return c.ClientIP()
		}
	default:
		// log warning if unknown key is used
		mw.logger.Warn("unknown key, using default",
			zap.String("key", string(key)),
			zap.String("default", string(ClientIP)),
		)

		return func(c *gin.Context) string {
			return c.ClientIP()
		}
	}
}
