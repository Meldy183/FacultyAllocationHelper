package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"gitlab.pg.innopolis.university/f.markin/fah/Fah-parsing_demo/internal/infrastructure/logger"
	"go.uber.org/zap"
)

func LoggerMiddleware(baseLogger *logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Generate request ID
			requestID := generateRequestID()

			// Add request ID to context
			ctx := context.WithValue(r.Context(), logger.RequestIDKey, requestID)

			// Add logger to context
			ctx = context.WithValue(ctx, logger.LoggerKey, baseLogger)

			// Update request with new context
			r = r.WithContext(ctx)

			// Log request start
			baseLogger.Info(ctx, "HTTP request started",
				zap.String("method", r.Method),
				zap.String("Host", r.Host),
				zap.String("path", r.URL.Path),
				zap.String("remote_addr", r.RemoteAddr),
				zap.String("user_agent", r.UserAgent()),
			)

			// Wrap response writer to capture status code
			wrapped := &responseWriter{ResponseWriter: w, statusCode: 200}

			// Call next handler
			next.ServeHTTP(wrapped, r)

			// Log request completion
			baseLogger.Info(ctx, "HTTP request completed",
				zap.String("method", r.Method),
				zap.String("Host", r.Host),
				zap.String("path", r.URL.Path),
				zap.Int("status_code", wrapped.statusCode),
				zap.Duration("duration", time.Since(start)),
			)
		})
	}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func generateRequestID() string {
	return fmt.Sprintf("req-%d", time.Now().UnixNano())
}
