package logger

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LoggerKey    = "logger"
	RequestIDKey = "request_id"
)

type Logger struct {
	logger *zap.Logger
}

type LoggerConfig struct {
	Level          string `yaml:"level"`
	Format         string `yaml:"format"`
	Console        bool   `yaml:"console"`
	ServiceName    string `yaml:"service_name"`
	HTTPLogging    bool   `yaml:"http_logging"`
	LogRequestBody bool   `yaml:"log_request_body"`
	SlowRequestMs  int    `yaml:"slow_request_ms"`
}

func NewLogger(cfg LoggerConfig, ctx context.Context) (context.Context, error) {
	level, err := zapcore.ParseLevel(cfg.Level)
	if err != nil {
		return nil, fmt.Errorf("invalid log level %s: %w", cfg.Level, err)
	}
	var zapConfig zap.Config
	if cfg.Format == "json" {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
	}
	zapConfig.Level = zap.NewAtomicLevelAt(level)
	zapConfig.InitialFields = map[string]interface{}{
		"service": cfg.ServiceName,
	}
	if cfg.Console {
		zapConfig.OutputPaths = []string{"stdout"}
		zapConfig.ErrorOutputPaths = []string{"stderr"}
	}
	logger, err := zapConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build logger: %w", err)
	}
	ctx = context.WithValue(ctx, LoggerKey, &Logger{logger: logger})
	return ctx, nil
}
func GetFromContext(ctx context.Context) *Logger {
	if logger, ok := ctx.Value(LoggerKey).(*Logger); ok {
		return logger
	}
	return &Logger{logger: zap.NewNop()}
}
func (cfg LoggerConfig) GetSlowRequestThreshold() time.Duration {
	return time.Duration(cfg.SlowRequestMs) * time.Millisecond
}
func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = l.addContextFields(ctx, fields)
	l.logger.Info(msg, fields...)
}

func (l *Logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = l.addContextFields(ctx, fields)
	l.logger.Error(msg, fields...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	fields = l.addContextFields(ctx, fields)
	l.logger.Fatal(msg, fields...)
}
func (l *Logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	fields = l.addContextFields(ctx, fields)
	l.logger.Warn(msg, fields...)
}

// Helper method to avoid duplication
func (l *Logger) addContextFields(ctx context.Context, fields []zap.Field) []zap.Field {
	if requestID, ok := ctx.Value(RequestIDKey).(string); ok && requestID != "" {
		fields = append(fields, zap.String("request_id", requestID))
	}
	return fields
}
