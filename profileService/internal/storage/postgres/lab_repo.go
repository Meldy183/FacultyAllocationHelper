package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/lab"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ lab.Repository = (*LabRepo)(nil)

type LabRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewLabRepo(pool *pgxpool.Pool, logger *zap.Logger) *LabRepo {
	return &LabRepo{pool: pool, logger: logger}
}

const (
	queryGetAllLabs           = `SELECT lab_id FROM lab`
	queryGetLabsByInstituteID = `SELECT lab_id FROM lab WHERE institute_id = $1`
	queryGetLabByID           = `SELECT lab_id, name, institute_id FROM lab WHERE lab_id = $1`
)

func (r *LabRepo) GetLabByID(ctx context.Context, labID int64) (*lab.Lab, error) {
	var labByID lab.Lab
	err := r.pool.QueryRow(ctx, queryGetLabByID, labID).Scan(
		&labByID.LabID,
		&labByID.Name,
		&labByID.InstituteID,
	)
	if err != nil {
		r.logger.Error("failed to query lab by id",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetLabByID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to query lab by id: %w", err)
	}
	r.logger.Info("Lab successfully retrieved",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetLabByID),
	)
	return &labByID, err
}

func (r *LabRepo) GetAllLabs(ctx context.Context) ([]int64, error) {
	rows, err := r.pool.Query(ctx, queryGetAllLabs)
	var labs []int64
	if err != nil {
		r.logger.Error("get-all",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetAllLabs),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error in get-all: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Err(); err != nil {
			r.logger.Error("get-all",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllLabs),
				zap.Error(err),
			)
			return nil, fmt.Errorf("error in get-all: %w", err)
		}
		var labToAdd int64
		err := rows.Scan(
			&labToAdd,
		)
		if err != nil {
			r.logger.Error("get-all",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllLabs),
				zap.Error(err),
			)
			return nil, fmt.Errorf("error in get-all: %w", err)
		}
		labs = append(labs, labToAdd)
	}
	r.logger.Info("get-all Success",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetAllLabs),
		zap.Int("labs", len(labs)),
	)
	return labs, nil
}

func (r *LabRepo) GetLabsByInstituteID(ctx context.Context, instituteID int64) ([]int64, error) {
	rows, err := r.pool.Query(ctx, queryGetLabsByInstituteID, instituteID)
	var labs []int64
	if err != nil {
		r.logger.Error("get-labs-by-institute-id",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetLabsByInstituteID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error in get-labs-by-institute-id: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Err(); err != nil {
			r.logger.Error("get-labs-by-institute-id",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetLabsByInstituteID),
				zap.Error(err),
			)
			return nil, fmt.Errorf("error in get-labs-by-institute-id: %w", err)
		}
		var labToAdd int64
		err := rows.Scan(
			&labToAdd,
		)
		if err != nil {
			r.logger.Error("get-labs-by-institute-id",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetLabsByInstituteID),
				zap.Error(err),
			)
			return nil, fmt.Errorf("error in get-labs-by-institute-id: %w", err)
		}
		labs = append(labs, labToAdd)
	}
	r.logger.Info("get-labs-by-institute-id success",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetLabsByInstituteID),
		zap.Int("labs", len(labs)),
	)
	return labs, nil
}
