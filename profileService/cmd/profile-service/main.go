package main

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/config"
	userprofile2 "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/http"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/institute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/position"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileInstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileLanguage"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileVersion"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/workload"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/storage/db"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/storage/postgres"
	"go.uber.org/zap"
	httpNet "net/http"
	"time"
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
	time.Sleep(time.Second * 3)
	defer pool.Close()
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
	userProfileRepo := postgres.NewFacultyProfileRepo(pool, logger)
	userLanguageRepo := postgres.NewUserLanguageRepo(pool, logger)
	userInstituteRepo := postgres.NewUserInstituteRepo(pool, logger)
	userCourseInstanceRepo := postgres.NewUserCourseInstance(pool, logger)
	positionRepo := postgres.NewPositionRepo(pool, logger)
	instituteRepo := postgres.NewInstituteRepo(pool, logger)
	profileVersionRepo := postgres.NewUserProfileVersionRepo(pool, logger)
	workloadRepo := postgres.NewSemesterWorkloadRepo(pool, logger)
	// Service layer inits
	userProfileService := facultyProfile.NewService(userProfileRepo, logger)
	userLanguageService := profileLanguage.NewService(userLanguageRepo, logger)
	userInstituteService := profileInstitute.NewService(userInstituteRepo, logger)
	userCourseInstanceService := profileCourseInstance.NewService(userCourseInstanceRepo, logger)
	positionService := position.NewService(positionRepo, logger)
	instituteService := institute.NewService(instituteRepo, logger)
	profileVersionService := profileVersion.NewService(profileVersionRepo, logger)
	workloadService := workload.NewService(workloadRepo, logger)
	handler := userprofile2.NewHandler(
		userProfileService,
		userInstituteService,
		userLanguageService,
		userCourseInstanceService,
		positionService,
		instituteService,
		profileVersionService,
		workloadService,
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
	logger.Info("Started Server",
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
