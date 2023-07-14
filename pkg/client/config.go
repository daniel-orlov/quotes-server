package client

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config is a configuration for client.
type Config struct {
	// Logging is the logging configuration.
	Logging struct {
		// LogLevel is the log level to use.
		Level string `envconfig:"LOG_LEVEL" default:"debug"`
		// LogFormat is the log format to use.
		Format string `envconfig:"LOG_FORMAT" default:"console"`
	}

	// Connection is the connection configuration.
	Connection struct {
		// ServerHost is the host of server, that client will connect to
		ServerHost string `envconfig:"SERVER_HOST" default:"localhost"`
		// ServerPort is the port of server that the client will connect to
		ServerPort int `envconfig:"SERVER_PORT" default:"8080"`
		// RequestPath is the path of request that client will send to server
		RequestPath string `envconfig:"REQUEST_PATH" default:"/v1/quotes/random"`
		// RequestRatePerSecond defines how many requests per second client will send to server
		RequestRatePerSecond int `envconfig:"REQUEST_RATE_PER_SECOND" default:"100"`
		// RequestCount defines how many requests client will send to server.
		// If RequestCount is 0, client will send requests infinitely
		RequestCount int `envconfig:"REQUEST_COUNT" default:"0"`
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
