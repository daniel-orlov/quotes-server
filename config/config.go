package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
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
