package logger

import (
	"context"

	"go.uber.org/zap"
)

const (
	LoggerID = "logger"
)

func WithContext(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, LoggerID, logger)
}
func GetLoggerFromCtx(ctx context.Context) *zap.Logger {
	l, ok := ctx.Value(LoggerID).(*zap.Logger)
	if !ok {
		return zap.NewNop()
	}
	return l
}
