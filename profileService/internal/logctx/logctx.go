package logctx

import (
	"context"
	"go.uber.org/zap"
)

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, "logger", logger)
}

// Logger Extract logger from context
func Logger(ctx context.Context) *zap.Logger {
	l, ok := ctx.Value("logger").(*zap.Logger)
	if !ok {
		return zap.NewNop() // no-op logger to avoid nil pointer panics
	}
	return l
}
