package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/config"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/storage/db"
	"go.uber.org/zap"
)

func main() {
	logger, loggerErr := initLogger()
	if loggerErr != nil {
		panic(loggerErr)
	}
	defer logger.Sync()
	ctx := logctx.WithLogger(context.Background(), logger)
	if err := godotenv.Load(); err != nil {
		logger.Fatal("Error loading .env file", zap.Error(err))
	}
	cfg := config.MustLoadConfig()
	pool, err := db.NewPostgresPool(ctx, cfg.Database)
	if err != nil {
		logger.Fatal("Error connecting to database", zap.Error(err))
	}
	logger.Info(fmt.Sprintf("Connected to database %v", cfg.Database))
	defer pool.Close()
}

func initLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout"} // redirect logs to stdout
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
