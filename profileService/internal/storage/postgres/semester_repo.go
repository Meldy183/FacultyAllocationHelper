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
)

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
