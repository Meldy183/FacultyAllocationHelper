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
	queryProgramByID    = `SELECT program_id, name FROM program WHERE code = $1`
	queryGetAllPrograms = `SELECT program_id, name FROM program`
)

func (r *ProgramRepo) GetProgramNameByID(ctx context.Context, id int) (*string, error) {
	row := r.pool.QueryRow(ctx, queryProgramByID, id)
	var program program.Program
	err := row.Scan(&program.ProgramID, &program.Name)
	if err != nil {
		r.logger.Error("failed to Get Program Name By ID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetProgramNameByID),
			zap.Int("id", id),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetLanguageByCode: %w", err)
	}
	r.logger.Info("successfully Got ProgramName By ID",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetProgramNameByID),
		zap.Int("id", id),
		zap.String("name", program.Name),
	)
	return &program.Name, nil
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
		var program program.Program
		err := rows.Scan(
			&program.ProgramID,
			&program.Name)
		if err != nil {
			r.logger.Error("failed to query all programs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllPrograms),
				zap.Error(err),
			)
			return nil, fmt.Errorf("get all programs: %w", err)
		}
		programs = append(programs, &program)
	}
	r.logger.Info("all programs returned",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetAllPrograms),
		zap.Int("count", len(programs)),
	)
	return programs, nil
}
