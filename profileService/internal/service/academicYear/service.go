package academicYear

import (
	"context"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/academicYear"
	"go.uber.org/zap"
)

var _ academicYear.Repository = (*Service)(nil)

type Service struct {
	logger *zap.Logger
	repo   academicYear.Repository
}

func NewService(logger *zap.Logger, repo academicYear.Repository) *Service {
	return &Service{logger: logger, repo: repo}
}

func (s *Service) GetAcademicYearNameByID(ctx context.Context, yearID int64) (*string, error) {
	return s.repo.GetAcademicYearNameByID(ctx, yearID)
}
