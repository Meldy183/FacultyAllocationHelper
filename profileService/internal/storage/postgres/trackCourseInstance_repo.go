package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	trackcourseinstance "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/trackCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ trackcourseinstance.Repository = (*TrackCourseRepo)(nil)

type TrackCourseRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewTrackCourseRepo(pool *pgxpool.Pool, logger *zap.Logger) *TrackCourseRepo {
	return &TrackCourseRepo{pool: pool, logger: logger}
}

const (
	queryTrackCourseInstanceByInstanceID = `SELECT track_id FROM track_course_instance WHERE instance_id = $1`
	queryAddTracksToCourseInstance       = `INSERT INTO track_course_instance (track_id, instance_id) VALUES ($1, $2) RETURNING track_course_instance_id`
)

func (r *TrackCourseRepo) AddTracksToCourseInstance(ctx context.Context, instanceID uuid.UUID, tracksIDs uuid.UUID) error {
	err := r.pool.QueryRow(ctx, queryAddTracksToCourseInstance, tracksIDs, instanceID).Scan(&instanceID)
	if err != nil {
		r.logger.Error("failed to add track to course instance",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", "AddTrackToCourseInstance"),
			zap.Error(err),
		)
		return fmt.Errorf("failed to add track to course instance: %w", err)
	}
	return nil
}

func (r *TrackCourseRepo) GetTracksIDsOfCourseByInstanceID(ctx context.Context, id uuid.UUID) ([]uuid.UUID, error) {
	rows, err := r.pool.Query(ctx, queryTrackCourseInstanceByInstanceID, id)
	if err != nil {
		r.logger.Error("failed to Get trackCourses By course ID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetTrackCourseByCourseID),
			zap.String("id", id.String()),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetTracksIDsOfCourseByInstanceID: %w", err)
	}
	defer rows.Close()
	var instances []uuid.UUID
	for rows.Next() {
		var trackID uuid.UUID
		err := rows.Scan(
			&trackID,
		)
		if err != nil {
			r.logger.Error("Error getting trackCourses by courseIDs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetTrackCourseByCourseID),
				zap.String("course id", id.String()),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetTracksIDsOfCourseByInstanceID failed: %w", err)
		}
		instances = append(instances, trackID)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error getting trackCourses by courseIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetTrackCourseByCourseID),
			zap.String("course id", id.String()),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetTracksIDsOfCourseByInstanceID failed: %w", err)
	}
	r.logger.Info("tracks found by course id",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetTrackCourseByCourseID),
		zap.Int("instancesLen", len(instances)),
	)
	return instances, nil
}
