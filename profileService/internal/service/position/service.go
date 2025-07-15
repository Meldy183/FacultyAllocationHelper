package position

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/position"
	"go.uber.org/zap"
)

var _ position.Repository = (*Service)(nil)

type Service struct {
	repo   position.Repository
	logger *zap.Logger
}

func NewService(repo position.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetPositionByID(ctx context.Context, positionID int) (*string, error) {
	if positionID <= 0 || positionID > 7 {
		s.logger.Error("position_id is invalid",
			zap.String("layer", logctx.LogGetPositionByID),
			zap.String("function", logctx.LogGetPositionByID),
			zap.Int("position_id", positionID),
		)
		return nil, fmt.Errorf("invalid position_id: %d", positionID)
	}
	positionByID, err := s.repo.GetPositionByID(ctx, positionID)
	if err != nil {
		s.logger.Error("failed to retrieve position by ID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetPositionByID),
			zap.Int("position_id", positionID),
			zap.Error(err),
		)
		return nil, err
	}
	s.logger.Info("Successfully retrieved position: ",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetPositionByID),
		zap.Int("position_id:", positionID),
	)
	return positionByID, nil
}

func (s *Service) GetAllPositions(ctx context.Context) ([]*position.Position, error) {
	positions, err := s.repo.GetAllPositions(ctx)
	if err != nil {
		s.logger.Error("failed to get all positions",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetAllPositions),
			zap.Error(err),
		)
		return nil, err
	}
	s.logger.Info("Successfully got all",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetAllPositions),
		zap.Int("positions", len(positions)),
	)
	return positions, nil
}
