package main

import (
	"context"
	"fmt"
	httpNet "net/http"
	"time"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/config"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/courses"
	userprofile2 "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/filters"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/parse"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/http"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/academicYear"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/completeCourse"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/completeUser"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/course"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/courseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/institute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/language"
	Parsing "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/parsing"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/position"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileInstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileLanguage"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileVersion"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/program"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/programCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/responsibleInstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/semester"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/staff"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/track"
	trackcourseinstance "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/trackCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/workload"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/storage/db"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/storage/postgres"
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
	profileRepo := postgres.NewFacultyProfileRepo(pool, logger)
	profileLanguageRepo := postgres.NewUserLanguageRepo(pool, logger)
	profileInstituteRepo := postgres.NewUserInstituteRepo(pool, logger)
	profileCourseInstanceRepo := postgres.NewUserCourseInstance(pool, logger)
	positionRepo := postgres.NewPositionRepo(pool, logger)
	languageRepo := postgres.NewLanguageRepo(pool, logger)
	instituteRepo := postgres.NewInstituteRepo(pool, logger)
	profileVersionRepo := postgres.NewUserProfileVersionRepo(pool, logger)
	workloadRepo := postgres.NewSemesterWorkloadRepo(pool, logger)
	programRepo := postgres.NewProgramRepo(pool, logger)
	staffRepo := postgres.NewStaffRepo(pool, logger)
	trackRepo := postgres.NewTrackRepo(pool, logger)
	courseRepo := postgres.NewCourseRepo(pool, logger)
	semesterRepo := postgres.NewSemesterRepo(pool, logger)
	academicYearRepo := postgres.NewAcademicYearRepo(pool, logger)
	responsibleInstituteRepo := postgres.NewResponsibleInstituteRepo(pool, logger)
	courseInstanceRepo := postgres.NewCourseInstanceRepo(pool, logger)
	trackInstanceRepo := postgres.NewTrackCourseRepo(pool, logger)
	programCourseInstanceRepo := postgres.NewProgramCourseRepo(pool, logger)
	// Service layer inits
	profileService := facultyProfile.NewService(profileRepo, logger)
	profileLanguageService := profileLanguage.NewService(profileLanguageRepo, logger)
	userCourseInstanceService := profileCourseInstance.NewService(profileCourseInstanceRepo, logger)
	positionService := position.NewService(positionRepo, logger)
	profileInstituteService := profileInstitute.NewService(profileInstituteRepo, logger)
	languageService := language.NewService(languageRepo, logger)
	instituteService := institute.NewService(instituteRepo, logger)
	profileVersionService := profileVersion.NewService(profileVersionRepo, logger)
	workloadService := workload.NewService(workloadRepo, logger)
	programService := program.NewService(programRepo, programCourseInstanceRepo, logger)
	trackService := track.NewService(trackRepo, trackInstanceRepo, logger)
	trackInstanceService := trackcourseinstance.NewService(trackInstanceRepo, logger)
	programCourseInstanceService := programCourseInstance.NewService(programCourseInstanceRepo, logger)
	courseInstanceService := courseInstance.NewService(courseInstanceRepo, logger)
	courseService := course.NewService(courseRepo, logger)
	fullCourseService := completeCourse.NewService(courseInstanceService, courseService, trackService, programService, trackInstanceService, programCourseInstanceService, logger)
	fullUserService := completeUser.NewService(logger, profileService, profileVersionService, languageService, profileLanguageService, instituteService, profileInstituteService)
	parseService := Parsing.NewService(logger, fullCourseService, fullUserService)
	staffService := staff.NewStaffService(staffRepo, logger)
	academicYearService := academicYear.NewService(academicYearRepo, logger)
	semesterService := semester.NewService(semesterRepo, logger)
	responsibleInstituteService := responsibleInstitute.NewService(responsibleInstituteRepo, logger)
	facultyHandler := userprofile2.NewHandler(
		profileService,
		profileInstituteRepo,
		profileLanguageService,
		userCourseInstanceService,
		positionService,
		instituteService,
		profileVersionService,
		*workloadService,
		logger,
	)
	courseHandler := courses.NewHandler(
		fullCourseService,
		staffService,
		profileInstituteService,
		*academicYearService,
		*semesterService,
		responsibleInstituteService,
		profileVersionService,
		profileService,
		courseInstanceService,
		courseService,
		programService,
		trackService,
		programCourseInstanceService,
		trackInstanceService,
		logger,
	)
	filtersHandler := filters.NewHandler(
		positionService,
		instituteService,
		academicYearService,
		programService,
		semesterService,
		responsibleInstituteService,
		logger,
	)
	parsingHandler := parse.NewHandler(
		logger,
		parseService,
	)
	router := http.NewRouter(facultyHandler, courseHandler, filtersHandler, parsingHandler)
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
	cfg.OutputPaths = []string{"stdout"}
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
