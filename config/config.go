package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"

	"github.com/daniel-orlov/quotes-server/internal/transport/middleware/ratelimiter"
)

// Config is the main configuration of the application.
// It is populated from environment variables and defaults.
type Config struct {
	// Logging is the logging configuration.
	Logging struct {
		// LogLevel is the log level to use.
		Level string `envconfig:"LOG_LEVEL" default:"debug"`
		// LogFormat is the log format to use.
		Format string `envconfig:"LOG_FORMAT" default:"console"`
	}
	// Server is the http server configuration.
	Server struct {
		// Port is the port to listen on.
		Port int `envconfig:"SERVER_PORT" default:"8080"`
		// Meddlewares is the configuration for the middlewares.
		Middlewares struct {
			// Ratelimiter is the configuration for the ratelimiter middleware.
			Ratelimiter struct {
				// Rate is the rate at which requests are allowed.
				Rate ratelimiter.Rate `envconfig:"RATELIMITER_RATE" default:"second"`
				// Limit is the maximum number of requests allowed.
				Limit uint `envconfig:"RATELIMITER_LIMIT" default:"5"`
				// Key is the key to use for the ratelimiter.
				Key ratelimiter.Key `envconfig:"RATELIMITER_KEY" default:"client_ip"`
			}
		}
	}
}

// NewConfig returns a new Config instance, populated with environment variables and defaults.
func NewConfig() (*Config, error) {
	// Create a new Config instance
	cfg := &Config{}

	// Populate the Config instance with environment variables
	err := envconfig.Process("", cfg)
	// Return an error if the environment variables could not be processed
	if err != nil {
		return nil, fmt.Errorf("processing environment variables: %w", err)
	}

	// Return the Config instance and no error
	return cfg, nil
}
