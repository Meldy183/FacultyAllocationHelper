package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/academicYear"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ academicYear.Repository = (*AcademicYearRepo)(nil)

type AcademicYearRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewAcademicYearRepo(pool *pgxpool.Pool, logger *zap.Logger) *AcademicYearRepo {
	return &AcademicYearRepo{
		pool:   pool,
		logger: logger,
	}
}

const (
	queryGetAcademicYearByID = `SELECT academic_year_name FROM academic_year WHERE academic_year_id = $1`
)

func (r *AcademicYearRepo) GetAcademicYearNameByID(ctx context.Context, yearID int64) (*string, error) {
	var str string
	err := r.pool.QueryRow(ctx, queryGetAcademicYearByID, yearID).Scan(
		&str,
	)
	if err != nil {
		r.logger.Error("Error getting academic year by id",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetAcademicYearNameByID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting academic year by id: %w", err)
	}
	return &str, nil
}
