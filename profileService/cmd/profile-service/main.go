package main

import (
	"context"
	"fmt"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/app"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/config"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/storage/db"
	"go.uber.org/zap"
)

func main() {
	logger, loggerErr := initLogger()
	if loggerErr != nil {
		fmt.Println("Panic occured")
		panic(loggerErr)
	}
	defer logger.Sync()
	ctx := context.Background()
	cfg := config.MustLoadConfig(logger)
	dataBase := db.NewConnectAndInit(logger)
	pool, err := dataBase.NewPostgresPool(ctx, cfg.Database)
	if err != nil {
		logger.Fatal("Error connecting to database",
			zap.String("layer", logctx.LogMainFuncLayer),
			zap.String("function", logctx.LogMain),
			zap.Error(err),
		)
	}
	logger.Info(fmt.Sprintf("Connection is completed  %v", cfg.Database),
		zap.String("layer", logctx.LogMainFuncLayer),
		zap.String("function", logctx.LogMain),
	)
	defer pool.Close()
	// time.Sleep(3 * time.Second)
	err = dataBase.InitSchema(ctx, pool)
	if err != nil {
		logger.Fatal("Error initializing schema",
			zap.String("layer", logctx.LogMainFuncLayer),
			zap.String("function", logctx.LogMain),
			zap.Error(err),
		)
	}
	logger.Info(
		"Schema initialized",
		zap.String("layer", logctx.LogMainFuncLayer),
		zap.String("function", logctx.LogMain),
	)

	// Repository layer inits
	app, appErr := app.InitializeApp(pool, logger)
	if appErr != nil {
		logger.Fatal("Failed to initialize app", zap.Error(appErr))
	}
	defer app.Close()

	err = app.Run(ctx, cfg.Server.Host, cfg.Server.Port, cfg.Server.ReadTimeout, cfg.Server.WriteTimeout, cfg.Server.IdleTimeout)
	if err != nil {
		logger.Fatal("Failed to run app",
			zap.String("layer", logctx.LogMainFuncLayer),
			zap.String("function", logctx.LogMain),
			zap.Error(err),
		)
	}
}

func initLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout"}
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
