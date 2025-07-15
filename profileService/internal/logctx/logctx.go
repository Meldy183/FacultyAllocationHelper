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
	LogMainFuncLayer           = "Main Function and Configuration"
	LogGetInstituteByID        = "getInstituteByID"
	LogGetAllInstitutes        = "getAllInstitutes"
	LogGetAllLabs              = "getAllLabs"
	LogGetLabsByInstituteID    = "getLabsByInstituteID"
	LogGetAllLanguages         = "getAllLanguages"
	LogGetLanguageByCode       = "getLanguageByCode"
	LogGetPositionByID         = "getPositionByID"
	LogGetAllPositions         = "getAllPositions"
	LogGetInstancesByProfileID = "getInstancesByProfileID"
	LogAddCourseInstance       = "addCourseInstance"
	LogGetInstituteByProfileID = "getInstituteByProfileID"
	LogAddUserInstitute        = "addUserInstitute"
	LogAddUserLanguage         = "addUserLanguage"
	LogGetLanguagesByProfileID = "getLanguagesByProfileID"
	LogAddProfile              = "addUserProfile"
	LogGetFacultyByProfileID   = "getFacultyByProfileID"
	LogUpdateFaculty           = "updateFaculty"
	LogGetFacultiesByFilters   = "getFacultiesByFilters"
	LogMustLoadConfig          = "mustLoadConfig"
)
