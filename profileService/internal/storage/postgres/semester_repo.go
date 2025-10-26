package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/semester"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ semester.Repository = (*SemesterRepo)(nil)

type SemesterRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewSemesterRepo(pool *pgxpool.Pool, logger *zap.Logger) *SemesterRepo {
	return &SemesterRepo{
		pool:   pool,
		logger: logger,
	}
}

const (
	queryGetSemesterNameByID = `SELECT semester_name FROM semester WHERE semester_id = $1`
	queryGetAllSemesters     = `SELECT semester_id, semester_name FROM semester`
)

func (r *SemesterRepo) GetAllSemesters(ctx context.Context) ([]semester.Semester, error) {
	rows, err := r.pool.Query(ctx, queryGetAllSemesters)
	if err != nil {
		r.logger.Error("Failed to get all semesters",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetAllSemesters),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to get all semesters: %w", err)
	}
	defer rows.Close()
	var semesters []semester.Semester
	for rows.Next() {
		var sem semester.Semester
		if err := rows.Scan(&sem.SemesterID, &sem.Name); err != nil {
			r.logger.Error("Failed to get all semesters",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllSemesters),
				zap.Error(err),
			)
			return nil, fmt.Errorf("failed to get all semesters: %w", err)
		}
		semesters = append(semesters, sem)
	}
	return semesters, nil
}

func (r *SemesterRepo) GetSemesterNameByID(ctx context.Context, semesterID int64) (*string, error) {
	var str string
	err := r.pool.QueryRow(ctx, queryGetSemesterNameByID, semesterID).Scan(&str)
	if err != nil {
		r.logger.Error("error getting semester name",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetSemesterNameByID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting semester name: %w", err)
	}
	return &str, nil
}
