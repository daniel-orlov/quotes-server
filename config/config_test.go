package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daniel-orlov/quotes-server/config"
)

func TestNewConfig_UsingDefaults(t *testing.T) {
	// Create a new Config instance
	cfg, err := config.NewConfig()

	// Assert that no error was returned
	assert.NoError(t, err)

	// Assert that the Config instance is not nil
	assert.NotNil(t, cfg)

	// Assert that the Config instance has the expected values
	assert.Equal(t, "debug", cfg.Logging.Level)
	assert.Equal(t, "console", cfg.Logging.Format)
	assert.Equal(t, 8080, cfg.Server.Port)
	assert.Equal(t, "second", cfg.Server.Middlewares.Ratelimiter.Rate)
	assert.Equal(t, 5, cfg.Server.Middlewares.Ratelimiter.Limit)
	assert.Equal(t, "client_ip", cfg.Server.Middlewares.Ratelimiter.Key)
}

func TestNewConfig_UsingEnvironmentVariables(t *testing.T) {
	// Set environment variables
	err := setEnvVars(map[string]string{
		"LOG_LEVEL":         "info",
		"LOG_FORMAT":        "json",
		"SERVER_PORT":       "80",
		"RATELIMITER_RATE":  "minute",
		"RATELIMITER_LIMIT": "10",
		"RATELIMITER_KEY":   "client_id",
	})
	// Assert that no error was returned
	assert.NoError(t, err)

	// Create a new Config instance
	cfg, err := config.NewConfig()

	// Assert that no error was returned
	assert.NoError(t, err)

	// Assert that the Config instance is not nil
	assert.NotNil(t, cfg)

	// Assert that the Config instance has the expected values
	assert.Equal(t, "info", cfg.Logging.Level)
	assert.Equal(t, "json", cfg.Logging.Format)
	assert.Equal(t, 80, cfg.Server.Port)
	assert.Equal(t, "minute", cfg.Server.Middlewares.Ratelimiter.Rate)
	assert.Equal(t, 10, cfg.Server.Middlewares.Ratelimiter.Limit)
	assert.Equal(t, "client_id", cfg.Server.Middlewares.Ratelimiter.Key)
}

// setEnvVars sets the given environment variables.
func setEnvVars(envVars map[string]string) error {
	// For each environment variable
	for key, value := range envVars {
		// Set the environment variable
		err := setEnvVar(key, value)
		// Return an error if the environment variable could not be set
		if err != nil {
			return err
		}
	}

	// Return no error
	return nil
}

// setEnvVar sets the given environment variable.
func setEnvVar(key, value string) error {
	// Set the environment variable using os.Setenv
	err := os.Setenv(key, value)
	// Return an error if the environment variable could not be set
	if err != nil {
		return fmt.Errorf("setting environment variable %s=%s: %w", key, value, err)
	}

	// Return no error
	return nil
}
