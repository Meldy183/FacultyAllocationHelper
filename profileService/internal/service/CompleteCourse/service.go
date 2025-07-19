package CompleteCourse

import (
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/CompleteCourse"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/course"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/courseInstance"
	"go.uber.org/zap"
)

var _ CompleteCourse.Service = (*Service)(nil)

type Service struct {
	logger          *zap.Logger
	instanceService *courseInstance.Service
	courseService   *course.Service
}

func NewService(instance *courseInstance.Service, course *course.Service, logger *zap.Logger) *Service {
	return &Service{
		instanceService: instance,
		courseService:   course,
		logger:          logger,
	}
}
