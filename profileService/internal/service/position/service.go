package position

import (
	"context"
	"fmt"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/position"
	"go.uber.org/zap"
)

var _ position.Repository = (*Service)(nil)

type Service struct {
	repo   position.Repository
	logger *zap.Logger
}

const (
	layer = "Service"
)

func NewService(repo position.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetByID(ctx context.Context, positionID int) (*string, error) {
	if positionID <= 0 || positionID > 7 {
		s.logger.Error("position_id is invalid")
		return nil, fmt.Errorf("invalid position_id: %d", positionID)
	}
	positionByID, err := s.repo.GetByID(ctx, positionID)
	if err != nil {
		s.logger.Error("failed to retrieve position by ID", zap.String("layer", layer),
			zap.Int("position_id", positionID), zap.Error(err))
		return nil, err
	}
	s.logger.Info("Successfully retrieved position: ",
		zap.String("LayerLevel", layer), zap.Int("position_id:", positionID))
	return positionByID, nil
}

func (s *Service) GetAll(ctx context.Context) ([]*position.Position, error) {
	positions, err := s.repo.GetAll(ctx)
	if err != nil {
		s.logger.Error("failed to get all positions", zap.String("layer", layer), zap.Error(err))
		return nil, err
	}
	s.logger.Info("Successfully got all", zap.String("layer", layer))
	return positions, nil
}
