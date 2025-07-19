package course

import (
	"context"
	"fmt"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/course"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ course.Service = (*Service)(nil)

type Service struct {
	logger *zap.Logger
	repo   course.Repository
}

func NewService(repo course.Repository, logger *zap.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) GetCourseByID(ctx context.Context, courseID int64) (*course.Course, error) {
	course, err := s.repo.GetCourseByID(ctx, courseID)
	if err != nil {
		s.logger.Error("error getting course by ID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetCourseByID),
			zap.Int64("courseID", courseID),
			zap.Error(err))
		return nil, fmt.Errorf("error getting course %w", err)
	}
	s.logger.Info("Course found",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetCourseByID),
		zap.Int64("courseID", courseID),
		zap.Any("course", course),
	)
	return course, nil
}
func (s *Service) AddCourse(ctx context.Context, course *course.Course) error {
	if responsibleInstituteIDValid(course.ResponsibleInstituteID) {
		s.logger.Error(
			"Invalid responsibleInstituteID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourse),
		)
		return fmt.Errorf("invalid responsibleInstituteID: %v", course.ResponsibleInstituteID)
	}
	err := s.repo.AddNewCourse(ctx, course)
	if err != nil {
		s.logger.Error(
			"Invalid responsibleInstituteID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourse),
		)
		return fmt.Errorf("error creating new course: %v", err)
	}
	s.logger.Info("new course created",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogAddNewCourse),
	)
	return nil
}

func responsibleInstituteIDValid(id int64) bool {
	return (id >= 1 && id <= 8)
}
