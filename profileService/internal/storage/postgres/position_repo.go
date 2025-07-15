package postgres

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/position"
	"go.uber.org/zap"
)

var _ position.Repository = (*PositionRepo)(nil)

type PositionRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewPositionRepo(pool *pgxpool.Pool, logger *zap.Logger) *PositionRepo {
	return &PositionRepo{pool: pool, logger: logger}
}

const (
	queryGetPositionByID = `SELECT position_id, name FROM position WHERE position_id = $1`
	queryGetAllPositions = `SELECT position_id, name FROM position`
)

func (r *PositionRepo) GetPositionByID(ctx context.Context, positionID int) (*string, error) {
	row := r.pool.QueryRow(ctx, queryGetPositionByID, positionID)
	var positionByID position.Position
	err := row.Scan(
		&positionByID.PositionID,
		&positionByID.Name)
	if err != nil {
		r.logger.Error("Error getting positionByID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetPositionByID),
			zap.Int("positionID", positionID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting positionByID: %w", err)
	}
	r.logger.Info("Successfully got positionByID",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetPositionByID),
		zap.Int("positionID", positionID),
	)
	return &positionByID.Name, nil
}

func (r *PositionRepo) GetAllPositions(ctx context.Context) ([]*position.Position, error) {
	r.logger.Info("Getting all positions")
	rows, err := r.pool.Query(ctx, queryGetAllPositions)
	if err != nil {
		r.logger.Error("Error getting all positions",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetAllPositions),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting all positions: %w", err)
	}
	defer rows.Close()
	var positions []*position.Position
	for rows.Next() {
		var iterThroughPositions position.Position
		err := rows.Scan(
			&iterThroughPositions.PositionID,
			&iterThroughPositions.Name)
		if err != nil {
			r.logger.Error("Error getting all positions",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllPositions),
				zap.Error(err),
			)
			return nil, fmt.Errorf("error getting all positions: %w", err)
		}
		positions = append(positions, &iterThroughPositions)
	}
	r.logger.Info("Finished getting all positions",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetAllPositions),
		zap.Int64("positions", int64(len(positions))),
	)
	return positions, nil
}
