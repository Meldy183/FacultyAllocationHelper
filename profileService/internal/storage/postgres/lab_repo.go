package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/lab"
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
	queryGetAllLabs           = `SELECT lab_id, name, institute_id FROM lab`
	queryGetLabsByInstituteID = `SELECT lab_id, name, institute_id FROM lab WHERE institute_id = $1`
	logGetAllLabs             = "GetAllLabs"
	logGetLabsByInstituteID   = "GetLabsByInstituteID"
)

func (r *LabRepo) GetAllLabs(ctx context.Context) ([]*lab.Lab, error) {
	rows, err := r.pool.Query(ctx, queryGetAllLabs)
	var labs []*lab.Lab
	if err != nil {
		r.logger.Error("get-all",
			zap.String("layer", logLayer),
			zap.String("function", logGetAllLabs),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error in get-all: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Err(); err != nil {
			r.logger.Error("get-all",
				zap.String("layer", logLayer),
				zap.String("function", logGetAllLabs),
				zap.Error(err),
			)
			return nil, fmt.Errorf("error in get-all: %w", err)
		}
		var labToAdd lab.Lab
		err := rows.Scan(
			&labToAdd.ID,
			&labToAdd.Name,
			&labToAdd.InstituteID,
		)
		if err != nil {
			r.logger.Error("get-all",
				zap.String("layer", logLayer),
				zap.String("function", logGetAllLabs),
				zap.Error(err),
			)
			return nil, fmt.Errorf("error in get-all: %w", err)
		}
		labs = append(labs, &labToAdd)
	}
	r.logger.Info("get-all Success",
		zap.String("layer", logLayer),
		zap.String("function", logGetAllLabs),
		zap.Int("labs", len(labs)),
	)
	return labs, nil
}

func (r *LabRepo) GetLabsByInstituteID(ctx context.Context, instituteID int64) ([]*lab.Lab, error) {
	rows, err := r.pool.Query(ctx, queryGetLabsByInstituteID, instituteID)
	var labs []*lab.Lab
	if err != nil {
		r.logger.Error("get-labs-by-institute-id",
			zap.String("layer", logLayer),
			zap.String("function", logGetLabsByInstituteID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error in get-labs-by-institute-id: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Err(); err != nil {
			r.logger.Error("get-labs-by-institute-id",
				zap.String("layer", logLayer),
				zap.String("function", logGetLabsByInstituteID),
				zap.Error(err),
			)
			return nil, fmt.Errorf("error in get-labs-by-institute-id: %w", err)
		}
		var labToAdd lab.Lab
		err := rows.Scan(
			&labToAdd.ID,
			&labToAdd.Name,
			&labToAdd.InstituteID)
		if err != nil {
			r.logger.Error("get-labs-by-institute-id",
				zap.String("layer", logLayer),
				zap.String("function", logGetLabsByInstituteID),
				zap.Error(err),
			)
			return nil, fmt.Errorf("error in get-labs-by-institute-id: %w", err)
		}
		labs = append(labs, &labToAdd)
	}
	r.logger.Info("get-labs-by-institute-id success",
		zap.String("layer", logLayer),
		zap.String("function", logGetLabsByInstituteID),
		zap.Int("labs", len(labs)),
	)
	return labs, nil
}
