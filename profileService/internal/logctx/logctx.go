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
	LogHandlerLayer                   = "Handler"
	LogServiceLayer                   = "Service"
	LogRepoLayer                      = "Repository"
	LogDBInitLayer                    = "DB Initialization"
	LogMainFuncLayer                  = "MainFunctionAndConfiguration"
	LogGetInstituteByID               = "getInstituteByID"
	LogGetAllInstitutes               = "getAllInstitutes"
	LogGetAllLabs                     = "getAllLabs"
	LogGetLabsByInstituteID           = "getLabsByInstituteID"
	LogGetAllLanguages                = "getAllLanguages"
	LogGetLanguageByCode              = "getLanguageByCode"
	LogGetPositionByID                = "getPositionByID"
	LogGetAllPositions                = "getAllPositions"
	LogAddCourseInstance              = "addCourseInstance"
	LogAddUserInstitute               = "addUserInstitute"
	LogAddUserLanguage                = "addUserLanguage"
	LogAddProfile                     = "addProfile"
	LogUpdateFaculty                  = "updateFaculty"
	LogGetUserLanguages               = "getUserLanguages"
	LogMustLoadConfig                 = "mustLoadConfig"
	LogMain                           = "Main"
	LogNewPostgresPool                = "NewPostgresPool"
	LogInitSchema                     = "InitSchema"
	LogGetProfileByID                 = "getProfile"
	LogGetAllFaculties                = "getAllFaculties"
	LogGetFacultyFilters              = "getFacultyFilters"
	LogGetUserInstitute               = "getUserInstitute"
	LogGetInstancesByProfileID        = "getInstancesByProfileID"
	LogGetProfileIDsByInstituteIDs    = "getProfileIDsByInstituteIDs"
	LogGetProfileIDsByPositionIDs     = "getProfileIDsByPositionIDs"
	LogGetVersionByProfileID          = "getVersionByProfileID"
	LogAddVersion                     = "addVersion"
	LogGetVersionIDByProfileID        = "getVersionIDByProfileID"
	LogAddProfileVersion              = "addProfileVersion"
	LogGetLogPages                    = "getLogPages"
	LogGetAcademicYearNameByID        = "getAcademicYearNameByID"
	LogGetSemesterWorkloadByVersionID = "getSemesterWorkloadByVersionID"
	LogAddSemesterWorkload            = "addSemesterWorkload"
	LogUpdateSemesterWorkload         = "updateSemesterWorkload"
	LogGetYearWorkloadByVersionID     = "getYearWorkloadByVersionID"
	LogGetStaffByInstanceID           = "getStaffByInstanceID"
)
