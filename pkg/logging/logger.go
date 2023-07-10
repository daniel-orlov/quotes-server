package logging

import (
	"fmt"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggerSingleton *zap.Logger
	once            sync.Once
)

// Logger returns a concurrence-safe singleton logger.
func Logger(format, level string) *zap.Logger {
	once.Do(func() {
		loggerSingleton = initLogger(format, level)
	})

	return loggerSingleton
}

func initLogger(format, level string) *zap.Logger {
	cfg := zap.Config{
		Encoding:         format,
		Level:            zap.NewAtomicLevelAt(levelFromString(level)),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "severity",
			TimeKey:        "timestamp",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	logger, err := cfg.Build(
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		panic(fmt.Errorf("creating logger: %w", err))
	}

	return logger
}

func levelFromString(level string) zapcore.Level {
	toLower := strings.ToLower(level)

	switch toLower {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
