package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/config"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/storage"
	"go.uber.org/zap"
	"log"
)

const (
	local = "local"
	prod  = "prod"
	dev   = "dev"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("starting server")
	ctx := context.Background()
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
