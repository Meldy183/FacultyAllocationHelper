package postgres

import (
	"context"
	"fmt"

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
	queryTrackCourseByID = `SELECT track_course_instance_id, track_id, instance_id FROM track_course_instance WHERE instance_id = $1`
)

func (r *TrackCourseRepo) GetTracksIDsOfCourseByInstanceID(ctx context.Context, id int) ([]*trackcourseinstance.TrackCourseInstance, error) {
	rows, err := r.pool.Query(ctx, queryTrackCourseByID, id)
	if err != nil {
		r.logger.Error("failed to Get trackCourses By course ID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetTrackCourseByCourseID),
			zap.Int("id", id),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetTracksIDsOfCourseByInstanceID: %w", err)
	}

	defer rows.Close()
	var instances []*trackcourseinstance.TrackCourseInstance
	for rows.Next() {
		var instance trackcourseinstance.TrackCourseInstance
		err := rows.Scan(&instance.TrackCourseInstanceID, &instance.TrackID, &instance.CourseInstanceID)
		if err != nil {
			r.logger.Error("Error getting trackCourses by courseIDs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetTrackCourseByCourseID),
				zap.String("course id", fmt.Sprintf("%v", id)),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetTracksIDsOfCourseByInstanceID failed: %w", err)
		}
		instances = append(instances, &instance)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error getting trackCourses by courseIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetTrackCourseByCourseID),
			zap.String("course id", fmt.Sprintf("%v", id)),
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
