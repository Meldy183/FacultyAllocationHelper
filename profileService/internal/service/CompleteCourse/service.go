package CompleteCourse

import (
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/courseInstance"
	"go.uber.org/zap"
)

type Service struct {
	logger          *zap.Logger
	instanceService *courseInstance.Service
	// TODO: add course service
}
