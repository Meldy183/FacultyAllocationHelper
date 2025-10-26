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
	courseObj, err := s.repo.GetCourseByID(ctx, courseID)
	if err != nil {
		s.logger.Error("error getting courseObj by ID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetCourseByID),
			zap.Int64("courseID", courseID),
			zap.Error(err))
		return nil, fmt.Errorf("error getting courseObj %w", err)
	}
	s.logger.Info("Course found",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetCourseByID),
		zap.Int64("courseID", courseID),
		zap.Any("courseObj", courseObj),
	)
	return courseObj, nil
}

func (s *Service) AddCourse(ctx context.Context, course *course.Course) error {
	if !responsibleInstituteIDValid(course.ResponsibleInstituteID) {
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
			"Error adding new course",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourse),
		)
		return fmt.Errorf("error creating new course: %w", err)
	}
	s.logger.Info("new course created",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogAddNewCourse),
	)
	return nil
}

func (s *Service) UpdateCourseByID(ctx context.Context, id int64, course *course.Course) error {
	if !responsibleInstituteIDValid(course.ResponsibleInstituteID) {
		s.logger.Error(
			"Invalid responsibleInstituteID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogUpdateCourseByID),
		)
		return fmt.Errorf("invalid responsibleInstituteID: %v", course.ResponsibleInstituteID)
	}
	err := s.repo.UpdateCourseByID(ctx, id, course)
	if err != nil {
		s.logger.Error(
			"Error updating course by id",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogUpdateCourseByID),
		)
		return fmt.Errorf("error creating new course: %w", err)
	}
	s.logger.Info("course updated",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogUpdateCourseByID),
	)
	return nil
}

func responsibleInstituteIDValid(id int64) bool {
	return id >= 1 && id <= 8
}
