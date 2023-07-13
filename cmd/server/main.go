package main

import (
	"fmt"
	"log"

	"go.uber.org/zap"

	"github.com/daniel-orlov/quotes-server/config"
	qsvc "github.com/daniel-orlov/quotes-server/internal/domain/service/quotes"
	cstore "github.com/daniel-orlov/quotes-server/internal/storage/challenges"
	qstore "github.com/daniel-orlov/quotes-server/internal/storage/quotes"
	httptransport "github.com/daniel-orlov/quotes-server/internal/transport/http"
	"github.com/daniel-orlov/quotes-server/internal/transport/http/quotes"
	"github.com/daniel-orlov/quotes-server/internal/transport/middleware/proofer"
	"github.com/daniel-orlov/quotes-server/internal/transport/middleware/ratelimiter"
	"github.com/daniel-orlov/quotes-server/pkg/logging"
	"github.com/daniel-orlov/quotes-server/pkg/pow"
)

func main() {
	// Parse config and check for errors.
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("parsing config: %v", err)
	}

	// Create a new logger.
	logger := logging.Logger(cfg.Logging.Format, cfg.Logging.Level)

	//--------------------------------------------------------------//
	//  				    	STORAGES                        	//
	//--------------------------------------------------------------//
	// Initialize the quote storage.
	quoteStorage := qstore.NewStorageInMemory(logger, qstore.GetQuotes())
	// Initialize the challenge storage.
	challengeStorage := cstore.NewStorageInMemory(logger)

	//--------------------------------------------------------------//
	//  				    	SERVICES                        	//
	//--------------------------------------------------------------//
	// Initialize the quote service.
	quoteService := qsvc.NewService(logger, quoteStorage)
	// Proof-of-work service.
	powService := pow.NewService(logger, challengeStorage)

	//--------------------------------------------------------------//
	//  				    	HANDLERS                        	//
	//--------------------------------------------------------------//
	// Initialize the quote handler.
	quotesHandler := quotes.NewHandler(logger, quoteService)

	//--------------------------------------------------------------//
	//  				    	MIDDLEWARES                     	//
	//--------------------------------------------------------------//
	// Rate limiter middleware.
	ratelimiterMW := ratelimiter.New(logger,
		&ratelimiter.Config{
			Rate:  cfg.Server.Middlewares.Ratelimiter.Rate,
			Limit: cfg.Server.Middlewares.Ratelimiter.Limit,
			Key:   cfg.Server.Middlewares.Ratelimiter.Key,
		})

	// Proof-of-work middleware.
	prooferMW := proofer.New(logger,
		&proofer.Config{
			ChallengeDifficulty: cfg.Server.Middlewares.Proofer.ChallengeDifficulty,
			SaltLength:          cfg.Server.Middlewares.Proofer.SaltLength,
		},
		powService,
	)

	// Initialize the Gin router.
	router := httptransport.NewRouter(quotesHandler, ratelimiterMW.Use(), prooferMW.Use())

	// Start the server and listen on port 8080.
	err = router.Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		logger.Fatal("running server failed", zap.Error(err))
	}
}
