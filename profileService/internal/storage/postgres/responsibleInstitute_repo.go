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
)

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
