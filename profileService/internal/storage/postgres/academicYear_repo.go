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
	queryAllAcademicYears    = `SELECT academic_year_id, academic_year_name FROM academic_year`
)

func (r *AcademicYearRepo) GetAllAcademicYears(ctx context.Context) ([]academicYear.AcademicYear, error) {
	rows, err := r.pool.Query(ctx, queryAllAcademicYears)
	if err != nil {
		r.logger.Error("failed to query all academic years",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetAllAcademicYears),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to query all academic years: %w", err)
	}
	defer rows.Close()
	var academicYears []academicYear.AcademicYear
	for rows.Next() {
		var academYear academicYear.AcademicYear
		if err := rows.Scan(&academYear.YearID, &academYear.Name); err != nil {
			r.logger.Error("failed to scan academic year name",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllAcademicYears),
				zap.Error(err),
			)
			return nil, fmt.Errorf("failed to scan academic year name: %w", err)
		}
		academicYears = append(academicYears, academYear)
	}
	return academicYears, nil
}

func (r *AcademicYearRepo) GetAcademicYearNameByID(ctx context.Context, yearID int64) (*string, error) {
	var str string
	r.logger.Info(queryGetAcademicYearByID, zap.Int64("year_id", yearID))
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
