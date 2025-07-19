package CompleteCourse

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/CompleteCourse"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/course"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/courseInstance"
	"go.uber.org/zap"
)

var _ CompleteCourse.Service = (*Service)(nil)

type Service struct {
	logger          *zap.Logger
	instanceService *courseInstance.Service
	courseService   *course.Service
	programsService *program
}

func NewService(instance *courseInstance.Service, course *course.Service, logger *zap.Logger) *Service {
	return &Service{
		instanceService: instance,
		courseService:   course,
		logger:          logger,
	}
}

func (s *Service) GetFullCourseInfoByID(ctx context.Context, instanceID int64) (*CompleteCourse.FullCourse, error) {
	courseInstanceObj, err := s.instanceService.GetCourseInstanceByID(ctx, instanceID)
	if err != nil {
		s.logger.Error("failed to fetch courseInstanceObj",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetFullCourseInfoByID),
			zap.Int64("instanceID", instanceID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to fetch courseInstanceObj: %w", err)
	}
	courseObj, err := s.courseService.GetCourseByID(ctx, courseInstanceObj.CourseID)
	if err != nil {
		s.logger.Error("failed to fetch courseObj",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetFullCourseInfoByID),
			zap.Int64("instanceID", instanceID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to fetch courseObj: %w", err)
	}
	fullCourse := &CompleteCourse.FullCourse{
		Course:         *courseObj,
		CourseInstance: *courseInstanceObj,
	}
	s.logger.Info("successfully fetched full course",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetFullCourseInfoByID),
		zap.Int64("instanceID", instanceID),
		zap.Int64("courseID", courseObj.CourseID),
		zap.Any("fullCourse", fullCourse),
	)
	return fullCourse, nil
}
