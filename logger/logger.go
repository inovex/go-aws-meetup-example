package logger

import (
	"context"
	"net/http"
	"time"

	"github.com/apex/gateway"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var _logger *zap.SugaredLogger

const ContextKey = iota

// init will be executed when the package is loaded
func init() {
	loggerConfig := zap.NewProductionConfig()
	logger, _ := loggerConfig.Build()
	_logger = logger.Sugar()
}

// EnableDevelopmentLogger constructs a development logger and overwrites the default production logger.
func EnableDevelopmentLogger() {
	loggerConfig := zap.NewDevelopmentConfig()
	logger, _ := loggerConfig.Build()
	_logger = logger.Sugar()
}

// Get returns the current default SugaredLogger.
func Get() *zap.SugaredLogger {
	return _logger
}

// FromRequest extracts the request scoped logger from the request.
// If no logger is found in the context, the default logger is returned.
func FromRequest(r *http.Request) *zap.SugaredLogger {
	return FromContext(r.Context())
}

// FromContext extracts the request scoped logger from the context.
// If no logger is found in the context, the default logger is returned.
func FromContext(ctx context.Context) *zap.SugaredLogger {
	l, ok := ctx.Value(ContextKey).(*zap.SugaredLogger)
	if !ok {
		return _logger
	}
	return l
}

// AddLoggerToContext returns a new context with the specified logger added.
// If ctx already contains a logger, it will be overwritten.
func AddLoggerToContext(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, ContextKey, logger)
}

func getRequestID(r *http.Request) string {
	ctx, ok := gateway.RequestContext(r.Context())
	if ok {
		// retrieve AWS request ID
		return ctx.RequestID
	} else {
		// generate request ID on our own
		return uuid.New().String()
	}
}

// AWSRequestIDMiddleware adds a request ID to the current logger.
func AWSRequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := getRequestID(r)
		ctx := AddLoggerToContext(r.Context(), _logger.With(
			"requestID", requestID,
			"userAgent", r.Header.Get("User-Agent"),
		))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequestDurationMiddleware times the request and logs its duration when finished
func RequestDurationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			diff := time.Since(start)
			_logger.Infow("request finished",
				"durationMillis", diff.Milliseconds(),
				"path", r.URL.Path,
			)
		}()
		next.ServeHTTP(w, r)
	})
}
