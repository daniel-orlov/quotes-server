// Package client contains the quotes server client, configured with the provided config.
package client

import (
	"net/http"

	"go.uber.org/zap"
)

// Client is a client for the quotes server.
type Client struct {
	logger *zap.Logger
	cfg    *Config
	client *http.Client
}

// NewClient creates a new quotes server client.
func NewClient(logger *zap.Logger, cfg *Config, client *http.Client) *Client {
	return &Client{logger: logger, cfg: cfg, client: client}
}
