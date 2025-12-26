package semester

import (
	"context"
	"errors"

	"github.com/google/uuid"
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

func (s *Service) GetSemesterNameByID(ctx context.Context, semesterID uuid.UUID) (*string, error) {
	if semesterID == uuid.Nil {
		s.logger.Error(`semesterID is invalid`,
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetSemesterNameByID),
		)
		return nil, errors.New("semesterID is invalid")
	}
	return s.repo.GetSemesterNameByID(ctx, semesterID)
}
