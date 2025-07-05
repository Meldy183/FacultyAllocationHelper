package postgres

import (
	"context"
	"fmt"
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
	queryGetByID = `SELECT institute_id, name FROM institute WHERE institute_id = $1`
	queryGetAll  = `SELECT institute_id, name FROM institute`
)

func (r *InstituteRepo) GetByID(ctx context.Context, instituteID int64) (*institute.Institute, error) {
	r.logger.Info("Getting instituteByID information by institute_id")
	row := r.pool.QueryRow(ctx, queryGetByID, instituteID)
	var instituteByID institute.Institute
	err := row.Scan(
		&instituteByID.InstituteID,
		&instituteByID.Name)
	if err != nil {
		r.logger.Error("Error getting instituteByID", zap.Error(err))
		return nil, fmt.Errorf("error getting instituteByID: %w", err)
	}
	r.logger.Info("Successfully got instituteByID", zap.Any("institute", instituteByID))
	return &instituteByID, nil
}

func (r *InstituteRepo) GetAll(ctx context.Context) ([]*institute.Institute, error) {
	r.logger.Info("Getting all institutes")
	rows, err := r.pool.Query(ctx, queryGetAll)
	if err != nil {
		r.logger.Error("Error getting all institutes", zap.Error(err))
		return nil, fmt.Errorf("error getting all institutes: %w", err)
	}
	var institutes []*institute.Institute
	for rows.Next() {
		var iterThroughInstitutes institute.Institute
		err := rows.Scan(
			&iterThroughInstitutes.InstituteID,
			&iterThroughInstitutes.Name)
		if err != nil {
			r.logger.Error("Error getting all institutes", zap.Error(err))
			return nil, fmt.Errorf("error getting all institutes: %w", err)
		}
		institutes = append(institutes, &iterThroughInstitutes)
	}
	r.logger.Info("Finished getting all institutes")
	return institutes, nil
}
