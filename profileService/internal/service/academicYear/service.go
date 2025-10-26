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

func (s *Service) GetAllAcademicYears(ctx context.Context) ([]academicYear.AcademicYear, error) {
	return s.repo.GetAllAcademicYears(ctx)
}

func NewService(repo academicYear.Repository, logger *zap.Logger) *Service {
	return &Service{logger: logger, repo: repo}
}

func (s *Service) GetAcademicYearNameByID(ctx context.Context, yearID int64) (*string, error) {
	s.logger.Warn("meow")
	return s.repo.GetAcademicYearNameByID(ctx, yearID)
}
