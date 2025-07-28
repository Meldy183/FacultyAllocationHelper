package postgres

import (
	"context"
	"fmt"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/institute"
	"go.uber.org/zap"
)

var _ institute.Repository = (*InstituteRepo)(nil)

type InstituteRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewInstituteRepo(pool *pgxpool.Pool, logger *zap.Logger) *InstituteRepo {
	return &InstituteRepo{pool: pool, logger: logger}
}

const (
	queryGetByID     = `SELECT institute_id, name FROM institute WHERE institute_id = $1`
	queryGetAll      = `SELECT institute_id, name FROM institute`
	queryGetIDByName = `SELECT institute_id FROM institute WHERE name = $1`
)

func (r *InstituteRepo) GetInstituteIDByName(ctx context.Context, instituteName string) (*int64, error) {
	row := r.pool.QueryRow(ctx, queryGetIDByName, instituteName)
	var instituteID int64
	err := row.Scan(&instituteID)
	if err != nil {
		r.logger.Error("Error getting instituteID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetInstituteIDByName),
			zap.String("instituteName", instituteName),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting instituteByID: %w", err)
	}
	r.logger.Info("Successfully got instituteID",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetInstituteIDByName),
		zap.Int64("instituteID", instituteID),
	)
	return &instituteID, nil
}
func (r *InstituteRepo) GetInstituteByID(ctx context.Context, instituteID int64) (*institute.Institute, error) {
	row := r.pool.QueryRow(ctx, queryGetByID, instituteID)
	var instituteByID institute.Institute
	err := row.Scan(
		&instituteByID.InstituteID,
		&instituteByID.Name)
	if err != nil {
		r.logger.Error("Error getting instituteByID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetInstituteByID),
			zap.Int64("instituteID", instituteID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting instituteByID: %w", err)
	}
	r.logger.Info("Successfully got instituteByID",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetInstituteByID),
		zap.Int64("instituteID", instituteID),
	)
	return &instituteByID, nil
}

func (r *InstituteRepo) GetAllInstitutes(ctx context.Context) ([]*institute.Institute, error) {
	rows, err := r.pool.Query(ctx, queryGetAll)
	if err != nil {
		r.logger.Error("Error getting all institutes",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetAllInstitutes),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting all institutes: %w", err)
	}
	defer rows.Close()
	var institutes []*institute.Institute
	for rows.Next() {
		var iterThroughInstitutes institute.Institute
		err := rows.Scan(
			&iterThroughInstitutes.InstituteID,
			&iterThroughInstitutes.Name)
		if err != nil {
			r.logger.Error("Error getting all institutes",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllInstitutes),
				zap.Error(err),
			)
			return nil, fmt.Errorf("error getting all institutes: %w", err)
		}
		institutes = append(institutes, &iterThroughInstitutes)
	}
	r.logger.Info("Finished getting all institutes",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetAllInstitutes),
		zap.Int64("institutes", int64(len(institutes))),
	)
	return institutes, nil
}
