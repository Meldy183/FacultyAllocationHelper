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
	queryGetByID        = `SELECT institute_id, name FROM institute WHERE institute_id = $1`
	queryGetAll         = `SELECT institute_id, name FROM institute`
	logGetInstituteByID = "GetInstituteByID"
	logGetAllInstitutes = "GetAllInstitutes"
)

func (r *InstituteRepo) GetInstituteByID(ctx context.Context, instituteID int64) (*institute.Institute, error) {
	row := r.pool.QueryRow(ctx, queryGetByID, instituteID)
	var instituteByID institute.Institute
	err := row.Scan(
		&instituteByID.InstituteID,
		&instituteByID.Name)
	if err != nil {
		r.logger.Error("Error getting instituteByID",
			zap.String("layer", logLayer),
			zap.String("function", logGetInstituteByID),
			zap.Int64("instituteID", instituteID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting instituteByID: %w", err)
	}
	r.logger.Info("Successfully got instituteByID",
		zap.String("layer", logLayer),
		zap.String("function", logGetInstituteByID),
		zap.Int64("instituteID", instituteID),
	)
	return &instituteByID, nil
}

func (r *InstituteRepo) GetAllInstitutes(ctx context.Context) ([]*institute.Institute, error) {
	rows, err := r.pool.Query(ctx, queryGetAll)
	if err != nil {
		r.logger.Error("Error getting all institutes",
			zap.String("layer", logLayer),
			zap.String("function", logGetAll),
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
				zap.String("layer", logLayer),
				zap.String("function", logGetAll),
				zap.Error(err),
			)
			return nil, fmt.Errorf("error getting all institutes: %w", err)
		}
		institutes = append(institutes, &iterThroughInstitutes)
	}
	r.logger.Info("Finished getting all institutes",
		zap.String("layer", logLayer),
		zap.String("function", logGetAll),
		zap.Int64("institutes", int64(len(institutes))),
	)
	return institutes, nil
}
