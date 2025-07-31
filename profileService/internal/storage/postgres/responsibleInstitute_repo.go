package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/responsibleInstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ responsibleInstitute.Repository = (*ResponsibleInstituteRepo)(nil)

type ResponsibleInstituteRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewResponsibleInstituteRepo(pool *pgxpool.Pool, logger *zap.Logger) *ResponsibleInstituteRepo {
	return &ResponsibleInstituteRepo{
		pool:   pool,
		logger: logger,
	}
}

const (
	queryGetResponsibleInstituteNameByID = `SELECT responsible_institute_name FROM
responsible_institute WHERE responsible_institute_id = $1`
	queryGetResponsibleInstituteIDByName = `SELECT responsible_institute_id FROM
	 responsible_institute WHERE responsible_institute_name = $1`
)

func (r *ResponsibleInstituteRepo) GetResponsibleInstituteIDByName(ctx context.Context, responsibleInstituteName string) (*int64, error) {
	ID := r.pool.QueryRow(ctx, queryGetResponsibleInstituteIDByName, responsibleInstituteName)
	var responsibleInstituteID int64
	err := ID.Scan(&responsibleInstituteID)
	if err != nil {
		r.logger.Error("Error getting responsible institute name",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetResponsibleInstituteIDByName),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting responsible institute name: %w", err)
	}
	r.logger.Info("successfully Got TrackName By ID",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetResponsibleInstituteIDByName),
		zap.String("name", responsibleInstituteName),
	)
	return &responsibleInstituteID, nil
}

func (r *ResponsibleInstituteRepo) GetAllInstitutes(ctx context.Context) ([]responsibleInstitute.ResponsibleInstitute, error) {
	rows, err := r.pool.Query(ctx, "SELECT * FROM responsible_institute")
	if err != nil {
		r.logger.Error("Error during query",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("layer", logctx.LogGetAllInstitutes),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error during query: %w", err)
	}
	defer rows.Close()
	var res []responsibleInstitute.ResponsibleInstitute
	for rows.Next() {
		var resp responsibleInstitute.ResponsibleInstitute
		err = rows.Scan(
			&resp.ResponsibleInstituteID,
			&resp.Name,
		)
		if err != nil {
			r.logger.Error("Error during scan",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("layer", logctx.LogGetAllInstitutes),
				zap.Error(err),
			)
			return nil, fmt.Errorf("error during scan: %w", err)
		}
		res = append(res, resp)
	}
	return res, nil
}

func (r *ResponsibleInstituteRepo) GetResponsibleInstituteNameByID(ctx context.Context, responsibleInstituteID int64) (*string, error) {
	str := r.pool.QueryRow(ctx, queryGetResponsibleInstituteNameByID, responsibleInstituteID)
	var responsibleInstituteName string
	err := str.Scan(&responsibleInstituteName)
	if err != nil {
		r.logger.Error("Error getting responsible institute name",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetResponsibleInstituteNameByID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting responsible institute name: %w", err)
	}
	return &responsibleInstituteName, nil
}
