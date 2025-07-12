package main

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/config"
	userprofile2 "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/userprofile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/http"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/usercourseinstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/userinstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/userlanguage"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/userprofile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/storage/db"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/storage/postgres"
	"go.uber.org/zap"
	httpNet "net/http"
)

const (
	logMain = "Main"
)

func main() {
	logger, loggerErr := initLogger()
	if loggerErr != nil {
		fmt.Println("Panic occured")
		panic(loggerErr)
	}
	defer logger.Sync()
	ctx := context.Background()
	cfg := config.MustLoadConfig()
	pool, err := db.NewPostgresPool(ctx, cfg.Database)
	if err != nil {
		logger.Fatal("Error connecting to database", zap.Error(err))
	}
	logger.Info(fmt.Sprintf("Connection is completed  %v", cfg.Database))
	defer pool.Close()
	err = db.InitSchema(ctx, pool)
	if err != nil {
		logger.Fatal("Error initializing schema",
			zap.String("function", logMain),
			zap.Error(err),
		)
	}
	// Repository layer inits
	userProfileRepo := postgres.NewUserProfileRepo(pool, logger)
	userLanguageRepo := postgres.NewUserLanguageRepo(pool, logger)
	userInstituteRepo := postgres.NewUserInstituteRepo(pool, logger)
	userCourseInstanceRepo := postgres.NewUserCourseInstance(pool, logger)
	// TODO languageRepo := postgres.NewLanguageRepo(pool, logger)
	// TODO labRepo := postgres.NewLabRepo(pool, logger)
	// TODO instituteRepo := postgres.NewInstituteRepo(pool, logger)
	// Service layer inits
	userProfileService := userprofile.NewService(userProfileRepo, logger)
	userLanguageService := userlanguage.NewService(userLanguageRepo, logger)
	userInstituteService := userinstitute.NewService(userInstituteRepo, logger)
	userCourseInstanceService := usercourseinstance.NewService(userCourseInstanceRepo, logger)
	// TODO languageService := language.NewService(languageRepo, logger)
	// TODO labService := lab.NewService(labRepo, logger)
	// TODO instituteService := institute.NewService(instituteRepo, logger)
	handler := userprofile2.NewHandler(
		userProfileService,
		userInstituteService,
		userLanguageService,
		userCourseInstanceService,
		logger,
	)
	router := http.NewRouter(handler)
	server := httpNet.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}
	logger.Info("Starting Server",
		zap.String("address", server.Addr),
	)
	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("Error starting server", zap.Error(err))
	}

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
