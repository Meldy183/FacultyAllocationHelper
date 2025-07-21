package semester

import (
	"context"
	"errors"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/semester"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ semester.Repository = (*Service)(nil)

type Service struct {
	logger *zap.Logger
	repo   semester.Repository
}

func (s *Service) GetAllSemesters(ctx context.Context) ([]semester.Semester, error) {
	return s.repo.GetAllSemesters(ctx)
}

func NewService(repo semester.Repository, logger *zap.Logger) *Service {
	return &Service{logger: logger, repo: repo}
}

func (s *Service) GetSemesterNameByID(ctx context.Context, semesterID int64) (*string, error) {
	if semesterID <= 0 || semesterID > 3 {
		s.logger.Error(`semesterID should be between 0 and 3`,
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetSemesterNameByID),
		)
		return nil, errors.New("semesterID should be between 0 and 3")
	}
	return s.repo.GetSemesterNameByID(ctx, semesterID)
}
