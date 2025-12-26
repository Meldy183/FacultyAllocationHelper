package lab

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/lab"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ lab.Repository = (*Service)(nil)

type Service struct {
	repo   lab.Repository
	logger *zap.Logger
}

func NewService(repo lab.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetLabByID(ctx context.Context, labID uuid.UUID) (*lab.Lab, error) {
	labObj, err := s.repo.GetLabByID(ctx, labID)
	if err != nil {
		s.logger.Info("Error getting labObj",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetLabByID),
			zap.Error(err),
		)
	}
	return labObj, err
}

func (s *Service) GetLabsByInstituteID(ctx context.Context, instituteID uuid.UUID) ([]uuid.UUID, error) {
	labs, err := s.repo.GetLabsByInstituteID(ctx, instituteID)
	if err != nil {
		s.logger.Error("Error getting labs",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetLabsByInstituteID),
			zap.String("institute_id", instituteID.String()),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting labs: %w", err)
	}
	s.logger.Info("Labs found",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetLabsByInstituteID),
		zap.String("institute_id", instituteID.String()),
	)
	return labs, nil
}

func (s *Service) GetAllLabs(ctx context.Context) ([]uuid.UUID, error) {
	labs, err := s.repo.GetAllLabs(ctx)
	if err != nil {
		s.logger.Error("Error getting labs",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetAllLabs),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting labs: %w", err)
	}
	s.logger.Info("Found all labs",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetAllLabs),
	)
	return labs, nil
}
