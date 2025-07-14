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
	logHandlerLayer            = "Handler"
	logServiceLayer            = "Service"
	logRepoLayer               = "Repository"
	logDBInitLayer             = "DB Initialization"
	logGetInstituteByID        = "getInstituteByID"
	logGetAllInstitutes        = "getAllInstitutes"
	logGetAllLabs              = "getAllLabs"
	logGetLabsByInstituteID    = "getLabsByInstituteID"
	logGetAllLanguages         = "getAllLanguages"
	logGetLanguageByCode       = "getLanguageByCode"
	logGetPositionByID         = "getPositionByID"
	logGetAllPositions         = "getAllPositions"
	logGetInstancesByProfileID = "getInstancesByProfileID"
	logAddCourseInstance       = "addCourseInstance"
	logGetInstituteByProfileID = "getInstituteByProfileID"
	logAddUserInstitute        = "addUserInstitute"
	logAddUserLanguage         = "addUserLanguage"
	logGetLanguagesByProfileID = "getLanguagesByProfileID"
	logAddProfile              = "addUserProfile"
	logGetFacultyByProfileID   = "getFacultyByProfileID"
	logUpdateFaculty           = "updateFaculty"
	logGetFacultiesByFilters   = "getFacultiesByFilters"
)
