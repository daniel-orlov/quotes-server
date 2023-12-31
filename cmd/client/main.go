package main

import (
	"log"

	"github.com/ybbus/httpretry"
	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/pkg/client"
	"github.com/daniel-orlov/quotes-server/pkg/logging"
)

func main() {
	// Parse config and check for errors.
	cfg, err := client.NewConfig()
	if err != nil {
		log.Fatalf("parsing config: %v", err)
	}

	// Create a new logger.
	logger := logging.Logger(cfg.Logging.Format, cfg.Logging.Level)

	// Log the config.
	logger.Debug("config", zap.Any("config", cfg))

	// Sync the logger before exiting.
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			log.Fatalf("syncing logger: %v", err)
		}
	}(logger)

	// Create a new HTTP client.
	httpClient := httpretry.NewDefaultClient()

	// Create a new client.
	quotesClient := client.NewClient(logger, cfg, httpClient)

	// Log successful client creation.
	logger.Info("client created")

	// Run the client in a loop.
	quotesClient.Run()

	// Log successful client run.
	logger.Info("client run completed")
}
