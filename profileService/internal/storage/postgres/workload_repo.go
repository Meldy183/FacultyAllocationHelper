package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	workloadDomain "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/workload"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ workloadDomain.Repository = (*WorkloadRepo)(nil)

type WorkloadRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

const (
	queryGetSemesterWorkloadByVersionID = `SELECT workload_id, profile_version_id, semester_id, lectures_count, tutorials_count, labs_count, electives_count, rate
FROM workload
WHERE profile_version_id = $1 AND semester_id = $2`
	queryAddSemesterWorkload = `INSERT INTO workload
	(profile_version_id, semester_id, lectures_count, tutorials_count, labs_count, electives_count, rate)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING workload_id`
	// queryUpdateSemesterWorkload = ``.
)

func NewSemesterWorkloadRepo(pool *pgxpool.Pool, log *zap.Logger) *WorkloadRepo {
	return &WorkloadRepo{pool: pool, logger: log}
}

func (r *WorkloadRepo) GetSemesterWorkloadByVersionID(
	ctx context.Context,
	VersionID int64,
	semID int64,
) (*workloadDomain.Workload, error) {
	row := r.pool.QueryRow(ctx, queryGetSemesterWorkloadByVersionID, VersionID, semID)
	var workload workloadDomain.Workload
	err := row.Scan(
		&workload.WorkloadID,
		&workload.ProfileVersionID,
		&workload.SemesterID,
		&workload.LecturesCount,
		&workload.TutorialsCount,
		&workload.LabsCount,
		&workload.ElectivesCount,
		&workload.Rate,
	)
	if err != nil {
		r.logger.Error("error getting workloadDomain from database",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetSemesterWorkloadByVersionID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting workloadDomain from database: %w", err)
	}
	r.logger.Info("workloadDomain successfully retrieved from database",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetSemesterWorkloadByVersionID),
		zap.Int64("semester_id", semID),
		zap.Int64("lectures_count", workload.LecturesCount),
		zap.Int64("tutorials_count", workload.TutorialsCount),
		zap.Int64("labs_count", workload.LabsCount),
		zap.Int64("electives_count", workload.ElectivesCount),
	)
	return &workload, nil
}

func (r *WorkloadRepo) AddSemesterWorkload(ctx context.Context, workload *workloadDomain.Workload) error {
	err := r.pool.QueryRow(ctx, queryAddSemesterWorkload,
		workload.ProfileVersionID,
		workload.SemesterID,
		workload.LecturesCount,
		workload.TutorialsCount,
		workload.LabsCount,
		workload.ElectivesCount,
		&workload.Rate,
	).Scan(&workload.WorkloadID)
	if err != nil {
		r.logger.Error("error adding workloadDomain from database",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogAddSemesterWorkload),
			zap.Error(err),
		)
		return fmt.Errorf("error adding workloadDomain from database: %w", err)
	}
	r.logger.Info("workloadDomain successfully added from database",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogAddSemesterWorkload),
	)
	return nil
}

func (r *WorkloadRepo) UpdateSemesterWorkload(ctx context.Context, workload *workloadDomain.Workload) error {
	return nil
	// TODO: implement me
}
