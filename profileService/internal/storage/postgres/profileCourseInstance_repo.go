package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ profileCourseInstance.Repository = (*ProfileCourseInstanceRepo)(nil)

type ProfileCourseInstanceRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewUserCourseInstance(pool *pgxpool.Pool, logger *zap.Logger) *ProfileCourseInstanceRepo {
	return &ProfileCourseInstanceRepo{pool: pool, logger: logger}
}

const (
	queryGetInstancesByProfileID = `SELECT profile_course_id FROM profile_course_instance WHERE profile_version_id = $1`
	queryAddCourseInstance       = `INSERT INTO profile_course_instance (profile_version_id, profile_course_id) VALUES ($1, $2)`
)

func (r *ProfileCourseInstanceRepo) GetCourseInstancesByVersionID(
	ctx context.Context,
	ProfileVersionID uuid.UUID,
) ([]uuid.UUID, error) {
	rows, err := r.pool.Query(ctx, queryGetInstancesByProfileID, ProfileVersionID)
	if err != nil {
		r.logger.Error("GetCourseInstancesByProfileID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetInstancesByProfileID),
			zap.String("profileID", ProfileVersionID.String()),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetCourseInstancesByProfileID: %w", err)
	}
	defer rows.Close()
	var instances []uuid.UUID
	for rows.Next() {
		if rows.Err() != nil {
			r.logger.Error("GetCourseInstancesByProfileID",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetInstancesByProfileID),
				zap.String("profileID", ProfileVersionID.String()),
				zap.Error(rows.Err()),
			)
			return nil, fmt.Errorf("GetCourseInstancesByProfileID: %w", rows.Err())
		}
		var instanceTaken uuid.UUID
		err := rows.Scan(&instanceTaken)
		if err != nil {
			r.logger.Error("GetCourseInstancesByProfileID",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetInstancesByProfileID),
				zap.String("profileID", ProfileVersionID.String()),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetCourseInstancesByProfileID: %w", err)
		}
		instances = append(instances, instanceTaken)
	}
	r.logger.Info("GetCourseInstancesByProfileID Success",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetInstancesByProfileID),
		zap.String("profileID", ProfileVersionID.String()),
	)
	return instances, nil
}

func (r *ProfileCourseInstanceRepo) AddCourseInstance(ctx context.Context,
	profileVersionCourseInstance *profileCourseInstance.ProfileVersionCourseInstance) error {
	_, err := r.pool.Exec(
		ctx,
		queryAddCourseInstance,
		profileVersionCourseInstance.ProfileCourseID,
		profileVersionCourseInstance.CourseInstanceID,
	)
	if err != nil {
		r.logger.Error("AddCourseInstance",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogAddCourseInstance),
			zap.Error(err),
		)
		return fmt.Errorf("AddCourseInstance: %w", err)
	}
	r.logger.Info("AddCourseInstance Success",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogAddCourseInstance),
		zap.String("profileID", profileVersionCourseInstance.ProfileCourseID.String()),
	)
	return nil
}
