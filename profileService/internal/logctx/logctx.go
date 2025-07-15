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

const (
	LogHandlerLayer            = "Handler"
	LogServiceLayer            = "Service"
	LogRepoLayer               = "Repository"
	LogDBInitLayer             = "DB Initialization"
	LogMainFuncLayer           = "MainFunctionAndConfiguration"
	LogGetInstituteByID        = "getInstituteByID"
	LogGetAllInstitutes        = "getAllInstitutes"
	LogGetAllLabs              = "getAllLabs"
	LogGetLabsByInstituteID    = "getLabsByInstituteID"
	LogGetAllLanguages         = "getAllLanguages"
	LogGetLanguageByCode       = "getLanguageByCode"
	LogGetPositionByID         = "getPositionByID"
	LogGetAllPositions         = "getAllPositions"
	LogAddCourseInstance       = "addCourseInstance"
	LogGetInstituteByProfileID = "getInstituteByProfileID"
	LogAddUserInstitute        = "addUserInstitute"
	LogAddUserLanguage         = "addUserLanguage"
	LogGetLanguagesByProfileID = "getLanguagesByProfileID"
	LogAddProfile              = "addProfile"
	LogGetFacultyByProfileID   = "getFacultyByProfileID"
	LogUpdateFaculty           = "updateFaculty"
	LogGetFacultiesByFilters   = "getFacultiesByFilters"
	LogGetUserLanguages        = "getUserLanguages"
	LogMustLoadConfig          = "mustLoadConfig"
	LogMain                    = "Main"
	LogNewPostgresPool         = "NewPostgresPool"
	LogInitSchema              = "InitSchema"
	LogGetProfile              = "getProfile"
	LogGetAllFaculties         = "getAllFaculties"
	LogGetFacultyFilters       = "getFacultyFilters"
	LogGetUserInstitute        = "getUserInstitute"
	LogGetProfilesByFilters    = "getProfilesByFilters"
)
