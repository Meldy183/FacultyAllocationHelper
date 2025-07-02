package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/config"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/storage"
	"go.uber.org/zap"
	"log"
)

func main() {
	logger, loggerErr := initLogger()
	if loggerErr != nil {
		panic(loggerErr)
	}
	defer logger.Sync()
	ctx := logctx.WithLogger(context.Background(), logger)
	if err := godotenv.Load(); err != nil {
		log.Fatalf("%s", err)
	}
	cfg := config.MustLoadConfig()
	pool, err := storage.NewPostgresPool(ctx, cfg.Database)
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	fmt.Println("connected to database")
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
