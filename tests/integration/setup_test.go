package integration_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

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

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

var (
	// testCfg is used to store the test config.
	testCfg *config.Config
	// testServer is used to mock HTTP responses.
	testServer *httptest.Server
	// testClient is used to mock HTTP requests.
	testClient *http.Client
	// testLogger is used to log test messages.
	testLogger *zap.Logger
)

// setup initializes the test server and client for integration testing.
func setup() {
	// Parse config and check for errors.
	var err error // this is needed to avoid shadowing the global testCfg variable.

	testCfg, err = config.NewConfig()
	if err != nil {
		log.Fatalf("parsing config: %v", err)
	}

	// Set rate limit to high otherwise tests will fail.
	testCfg.Server.Middlewares.Ratelimiter.Limit = uint(100)

	// Create a new logger.
	testLogger = logging.Logger(testCfg.Logging.Format, testCfg.Logging.Level)

	// Initialize the quote storage.
	quoteStorage := qstore.NewStorageInMemory(testLogger, qstore.GetQuotes())
	// Initialize the challenge storage.
	challengeStorage := cstore.NewStorageInMemory(testLogger)

	// Initialize the quote service.
	quoteService := qsvc.NewService(testLogger, quoteStorage)
	// Proof-of-work service.
	powService := pow.NewService(testLogger, challengeStorage)

	// Initialize the quote handler.
	quotesHandler := quotes.NewHandler(testLogger, quoteService)

	// Rate limiter middleware.
	ratelimiterMW := ratelimiter.New(testLogger,
		&ratelimiter.Config{
			Rate:  testCfg.Server.Middlewares.Ratelimiter.Rate,
			Limit: testCfg.Server.Middlewares.Ratelimiter.Limit,
			Key:   testCfg.Server.Middlewares.Ratelimiter.Key,
		})

	// Proof-of-work middleware.
	prooferMW := proofer.New(testLogger,
		&proofer.Config{
			ChallengeDifficulty: testCfg.Server.Middlewares.Proofer.ChallengeDifficulty,
			SaltLength:          testCfg.Server.Middlewares.Proofer.SaltLength,
		},
		powService,
	)

	// Initialize the Gin router.
	testRouter := httptransport.NewRouter(quotesHandler, ratelimiterMW.Use(), prooferMW.Use())

	// Create the test server.
	testServer = httptest.NewServer(testRouter)

	// Create a test HTTP client using the test server's URL.
	testClient = testServer.Client()

	// Log successful setup.
	testLogger.Info("integration test setup completed")
}

// Shutdown performs cleanup after the tests.
// It closes the test server and syncs the test logger.
// Since all the storages are in-memory, there is no need to purge them.
func shutdown() {
	// Close the test server.
	testServer.Close()

	// Sync the test logger.
	err := testLogger.Sync()
	if err != nil {
		log.Printf("syncing logger: %v", err)
	}
}
