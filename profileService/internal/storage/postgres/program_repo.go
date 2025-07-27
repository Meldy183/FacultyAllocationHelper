package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/program"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ program.Repository = (*ProgramRepo)(nil)

type ProgramRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewProgramRepo(pool *pgxpool.Pool, logger *zap.Logger) *ProgramRepo {
	return &ProgramRepo{pool: pool, logger: logger}
}

const (
	queryProgramByID     = `SELECT program_id, name FROM program WHERE program_id = $1`
	queryGetAllPrograms  = `SELECT program_id, name FROM program`
	queryProgramIDByName = `SELECT program_id FROM program WHERE name = $1`
)

func (r *ProgramRepo) GetProgramIDByName(ctx context.Context, name string) (*int64, error) {
	ID := r.pool.QueryRow(ctx, queryProgramByID, name)
	var ProgramID int64
	err := ID.Scan(&ProgramID)
	if err != nil {
		r.logger.Error("Error getting responsible institute name",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetProgramIDByName),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting responsible institute name: %w", err)
	}
	r.logger.Info("successfully Got ProgramName By ID",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetProgramIDByName),
		zap.String("name", name),
	)
	return &ProgramID, nil
}
func (r *ProgramRepo) GetProgramNameByID(ctx context.Context, id int64) (*string, error) {
	row := r.pool.QueryRow(ctx, queryProgramByID, id)
	var programObj program.Program
	err := row.Scan(&programObj.ProgramID, &programObj.Name)
	if err != nil {
		r.logger.Error("failed to Get Program Name By ID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetProgramNameByID),
			zap.Int64("id", id),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetProgramNameByCode: %w", err)
	}
	r.logger.Info("successfully Got ProgramName By ID",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetProgramNameByID),
		zap.Int64("id", id),
		zap.String("name", programObj.Name),
	)
	return &programObj.Name, nil
}
func (r *ProgramRepo) GetAllPrograms(ctx context.Context) ([]*program.Program, error) {
	rows, err := r.pool.Query(ctx, queryGetAllPrograms)
	if err != nil {
		r.logger.Error("failed to query all programs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetAllPrograms),
			zap.Error(err),
		)
		return nil, fmt.Errorf("get all programs: %w", err)
	}
	defer rows.Close()
	var programs []*program.Program
	for rows.Next() {
		if rows.Err() != nil {
			r.logger.Error("failed to query all programs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllPrograms),
				zap.Error(rows.Err()),
			)
			return nil, fmt.Errorf("get all programs: %w", rows.Err())
		}
		var programObj program.Program
		err := rows.Scan(
			&programObj.ProgramID,
			&programObj.Name)
		if err != nil {
			r.logger.Error("failed to query all programs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllPrograms),
				zap.Error(err),
			)
			return nil, fmt.Errorf("get all programs: %w", err)
		}
		programs = append(programs, &programObj)
	}
	r.logger.Info("all programs returned",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetAllPrograms),
		zap.Int("count", len(programs)),
	)
	return programs, nil
}
