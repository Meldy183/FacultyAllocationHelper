package position

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/position"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ position.Service = (*Service)(nil)

type Service struct {
	repo   position.Repository
	logger *zap.Logger
}

func NewService(repo position.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetPositionIDByName(ctx context.Context, positionName string) (*uuid.UUID, error) {
	positionID, err := s.repo.GetPositionIDByName(ctx, positionName)
	if err != nil || positionID == nil {
		s.logger.Error("failed to retrieve positionID by Name",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetPositionByID),
			zap.String("positionName", positionName),
			zap.Error(err),
		)
		return nil, err
	}
	s.logger.Info("Successfully retrieved positionID: ",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetPositionIDByName),
		zap.String("positionName", positionName),
	)
	return positionID, nil
}

func (s *Service) GetPositionByID(ctx context.Context, positionID uuid.UUID) (*string, error) {
	if positionID == uuid.Nil {
		s.logger.Error("position_id is invalid",
			zap.String("layer", logctx.LogGetPositionByID),
			zap.String("function", logctx.LogGetPositionByID),
			zap.String("position_id", positionID.String()),
		)
		return nil, fmt.Errorf("invalid position_id: %s", positionID.String())
	}
	positionByID, err := s.repo.GetPositionByID(ctx, positionID)
	if err != nil {
		s.logger.Error("failed to retrieve position by LabID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetPositionByID),
			zap.String("position_id", positionID.String()),
			zap.Error(err),
		)
		return nil, err
	}
	s.logger.Info("Successfully retrieved position: ",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetPositionByID),
		zap.String("position_id", positionID.String()),
	)
	return positionByID, nil
}

func (s *Service) GetAllPositions(ctx context.Context) ([]uuid.UUID, error) {
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
