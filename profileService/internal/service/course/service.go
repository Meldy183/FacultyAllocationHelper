package course

import (
	"context"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/course"
	"go.uber.org/zap"
)

var _ course.Service = (*Service)(nil)

type Service struct {
	logger *zap.Logger
	repo   *course.Repository
}

func NewService(repo *course.Repository, logger *zap.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) GetCourseByID(ctx context.Context, courseID int64) (*course.Course, error) {}
func (s *Service) AddCourse(ctx context.Context, course *course.Course) error                {}
