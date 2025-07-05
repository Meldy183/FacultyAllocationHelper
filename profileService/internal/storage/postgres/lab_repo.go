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

func NewLabRepo(logger *zap.Logger, pool *pgxpool.Pool) *LabRepo {
	return &LabRepo{pool: pool, logger: logger}
}

const (
	queryGetAllLabs           = `SELECT lab_id, name, institute_id FROM lab`
	queryGetLabsByInstituteID = `SELECT lab_id, name, institute_id FROM lab WHERE institute_id = $1`
)

func (r *LabRepo) GetAllLabs(ctx context.Context) ([]*lab.Lab, error) {
	r.logger.Info("get-all")
	rows, err := r.pool.Query(ctx, queryGetAllLabs)
	var labs []*lab.Lab
	if err != nil {
		r.logger.Error("get-all", zap.Error(err))
		return nil, fmt.Errorf("error in get-all: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Err(); err != nil {
			r.logger.Error("get-all", zap.Error(err))
			return nil, fmt.Errorf("error in get-all: %w", err)
		}
		var labToAdd lab.Lab
		err := rows.Scan(
			&labToAdd.ID,
			&labToAdd.Name,
			&labToAdd.InstituteID,
		)
		if err != nil {
			r.logger.Error("get-all", zap.Error(err))
			return nil, fmt.Errorf("error in get-all: %w", err)
		}
		labs = append(labs, &labToAdd)
	}
	r.logger.Info("get-all Success", zap.Int("labs", len(labs)))
	return labs, nil
}

func (r *LabRepo) GetLabsByInstituteID(ctx context.Context, instituteID int64) ([]*lab.Lab, error) {
	r.logger.Info("get-labs-by-institute-id")
	rows, err := r.pool.Query(ctx, queryGetLabsByInstituteID, instituteID)
	var labs []*lab.Lab
	if err != nil {
		r.logger.Error("get-labs-by-institute-id", zap.Error(err))
		return nil, fmt.Errorf("error in get-labs-by-institute-id: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Err(); err != nil {
			r.logger.Error("get-labs-by-institute-id", zap.Error(err))
			return nil, fmt.Errorf("error in get-labs-by-institute-id: %w", err)
		}
		var labToAdd lab.Lab
		err := rows.Scan(
			&labToAdd.ID,
			&labToAdd.Name,
			&labToAdd.InstituteID)
		if err != nil {
			r.logger.Error("get-labs-by-institute-id", zap.Error(err))
			return nil, fmt.Errorf("error in get-labs-by-institute-id: %w", err)
		}
		labs = append(labs, &labToAdd)
	}
	r.logger.Info("get-labs-by-institute-id success", zap.Int("labs", len(labs)))
	return labs, nil
}
