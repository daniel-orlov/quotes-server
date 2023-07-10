package main

import (
	"go.uber.org/zap"

	qsvc "github.com/daniel-orlov/quotes-server/internal/domain/service/quotes"
	qstore "github.com/daniel-orlov/quotes-server/internal/storage/quotes"
	httptransport "github.com/daniel-orlov/quotes-server/internal/transport/http"
	"github.com/daniel-orlov/quotes-server/internal/transport/http/quotes"
	"github.com/daniel-orlov/quotes-server/pkg/logging"
)

func main() {
	// Create a new logger
	logger := logging.Logger("json", "debug")

	// Initialize the quote storage.
	quotesStorage := qstore.NewStorageInMemory(logger, quoteDB)

	// Initialize the quote service.
	quoteService := qsvc.NewService(logger, quotesStorage)

	// Initialize the quote handler.
	quotesHandler := quotes.NewHandler(logger, quoteService)

	// Initialize the Gin router.
	router := httptransport.NewRouter(quotesHandler)

	// Start the server and listen on port 8080.
	err := router.Run(":8080")
	if err != nil {
		logger.Fatal("running server failed", zap.Error(err))
	}
}
