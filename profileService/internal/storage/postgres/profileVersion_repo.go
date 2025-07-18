package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ profileVersion.Repository = (*ProfileVersionRepo)(nil)

type ProfileVersionRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewUserProfileVersionRepo(pool *pgxpool.Pool, logger *zap.Logger) *ProfileVersionRepo {
	return &ProfileVersionRepo{pool: pool, logger: logger}
}

const (
	queryGetVersionByProfileID = `
		SELECT profile_version_id, profile_id, year, workload, maxload, position_id, employment_type, degree, mode 
		FROM user_profile_version
		WHERE profile_id = $1 AND year = $2
	`
	queryInsertVersion = `
		INSERT INTO user_profile_version (position_id, profile_id, year)
		VALUES ($1, $2, $3)
		RETURNING profile_version_id
	`
	queryUpdateVersion = `
		UPDATE user_profile_version
		SET profile_id = $1, year = $2, lectures_count = $3, tutorials_count = $4, labs_count = $5,
		elective_count = $6, workload = $7, maxload = $8, position_id = $9, employment_type = $10, degree = $11, mode = $12
		WHERE profile_version_id = $13
`
)

func (r *ProfileVersionRepo) GetVersionByProfileID(ctx context.Context, profileID int64, year int) (*profileVersion.ProfileVersion, error) {
	row := r.pool.QueryRow(ctx, queryGetVersionByProfileID, profileID, year)
	var version profileVersion.ProfileVersion
	err := row.Scan(
		&version.ProfileID,
		&version.ProfileID,
		&version.Year,
		&version.Workload,
		&version.MaxLoad,
		&version.PositionID,
		&version.EmploymentType,
		&version.Degree,
		&version.Mode,
	)
	if err != nil {
		r.logger.Error("Failed to get version by profile ID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetVersionByProfileID),
			zap.Int64("profile_id", profileID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to get version by profile ID: %w", err)
	}
	r.logger.Info("Succeeded to get version by profile ID",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetVersionByProfileID),
		zap.Int64("profile_id", profileID),
	)
	return &version, nil
}

func (r *ProfileVersionRepo) AddProfileVersion(ctx context.Context, version *profileVersion.ProfileVersion) error {
	err := r.pool.QueryRow(ctx, queryInsertVersion,
		version.PositionID,
		version.ProfileID,
		version.Year,
	).Scan(&version.ProfileVersionId)
	if err != nil {
		r.logger.Error("Failed to insert version",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogAddVersion),
			zap.Error(err),
		)
		return fmt.Errorf("insert version failed: %w", err)
	}
	r.logger.Info("Succeeded to insert version",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogAddVersion),
		zap.Int64("version_id", version.ProfileVersionId),
	)
	return nil
}
