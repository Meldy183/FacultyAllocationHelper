package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/position"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
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
	queryGetPositionByID = `SELECT name FROM position WHERE position_id = $1`
	queryGetAllPositions = `SELECT position_id FROM position`
)

func (r *PositionRepo) GetPositionByID(ctx context.Context, positionID int64) (*string, error) {
	row := r.pool.QueryRow(ctx, queryGetPositionByID, positionID)
	var posName string
	err := row.Scan(
		&posName,
	)
	if err != nil {
		r.logger.Error("Error getting posName",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetPositionByID),
			zap.Int64("positionID", positionID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting posName: %w", err)
	}
	r.logger.Info("Successfully got position by ID",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetPositionByID),
		zap.Int64("positionID", positionID),
	)
	return &posName, nil
}

func (r *PositionRepo) GetAllPositions(ctx context.Context) ([]int64, error) {
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
	var positions []int64
	for rows.Next() {
		var iterThroughPositions int64
		err := rows.Scan(
			&iterThroughPositions,
		)
		if err != nil {
			r.logger.Error("Error getting all positions",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllPositions),
				zap.Error(err),
			)
			return nil, fmt.Errorf("error getting all positions: %w", err)
		}
		positions = append(positions, iterThroughPositions)
	}
	r.logger.Info("Successfully got all positions",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetAllPositions),
		zap.Int64("positions", int64(len(positions))),
	)
	return positions, nil
}
