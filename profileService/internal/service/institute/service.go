package institute

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/institute"
	"go.uber.org/zap"
)

var _ institute.Repository = (*Service)(nil)

type Service struct {
	repo   institute.Repository
	logger *zap.Logger
}

const (
	layer = "Service"
)

func NewService(logger *zap.Logger, repo institute.Repository) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetByID(ctx context.Context, instituteID int64) (*institute.Institute, error) {
	if instituteID <= 0 {
		s.logger.Error("institute_id is invalid")
		return nil, fmt.Errorf("invalid institute_id: %d", instituteID)
	}
	instituteByID, err := s.repo.GetByID(ctx, instituteID)
	if err != nil {
		s.logger.Error("failed to retrieve institute by ID", zap.String("layer", layer),
			zap.Int64("institute_id", instituteID), zap.Error(err))
		return nil, err
	}
	s.logger.Info("Successfully retrieved institute: ",
		zap.String("LayerLevel", layer), zap.Int64("institute_id:", instituteID))
	return instituteByID, nil
}

func (s *Service) GetAll(ctx context.Context) ([]*institute.Institute, error) {
	institutes, err := s.repo.GetAll(ctx)
	if err != nil {
		s.logger.Error("failed to get all institutes", zap.String("layer", layer), zap.Error(err))
		return nil, err
	}
	s.logger.Info("Successfully got all", zap.String("layer", layer))
	return institutes, nil
}
