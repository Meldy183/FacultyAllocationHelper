package courseInstance

import (
	"context"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/courseInstance"
	"go.uber.org/zap"
)

var _ courseInstance.Repository = (*Service)(nil)

type Service struct {
	repo   courseInstance.Repository
	logger *zap.Logger
}

func NewService(repo courseInstance.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) AddNewCourseInstance(ctx context.Context, course *courseInstance.CourseInstance) error {

}

func (s *Service) GetCourseInstanceByID(ctx context.Context, id int64) (*courseInstance.CourseInstance, error) {

}

func (s *Service) UpdateCourseInstanceByID(ctx context.Context, course *courseInstance.CourseInstance) error {

}
