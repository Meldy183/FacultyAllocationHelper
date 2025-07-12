package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/usercourseinstance"
	"go.uber.org/zap"
)

var _ usercourseinstance.Repository = (*UserCourseInstanceRepo)(nil)

type UserCourseInstanceRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewUserCourseInstance(pool *pgxpool.Pool, logger *zap.Logger) *UserCourseInstanceRepo {
	return &UserCourseInstanceRepo{pool: pool, logger: logger}
}

const (
	queryGetInstancesByProfileID = `SELECT instance_id FROM user_course_instance WHERE profile_id = $1`
	queryAddCourseInstance       = `INSERT INTO user_course_instance (profile_id, instance_id) VALUES ($1, $2)`
	logGetInstancesByID          = "GetInstancesByID"
	logAddCourseInstance         = "AddCourseInstance"
)

func (r *UserCourseInstanceRepo) GetInstancesByProfileID(ctx context.Context, profileID int64) ([]int64, error) {
	rows, err := r.pool.Query(ctx, queryGetInstancesByProfileID, profileID)
	if err != nil {
		r.logger.Error("GetInstancesByProfileID",
			zap.String("layer", logLayer),
			zap.String("function", logGetUserInstituteByID),
			zap.Int64("profileID", profileID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetInstancesByProfileID: %w", err)
	}
	defer rows.Close()
	var instances []int64
	for rows.Next() {
		if rows.Err() != nil {
			r.logger.Error("GetInstancesByProfileID",
				zap.String("layer", logLayer),
				zap.String("function", logGetUserInstituteByID),
				zap.Int64("profileID", profileID),
				zap.Error(rows.Err()),
			)
			return nil, fmt.Errorf("GetInstancesByProfileID: %w", rows.Err())
		}
		var instanceTaken int64
		err := rows.Scan(&instanceTaken)
		if err != nil {
			r.logger.Error("GetInstancesByProfileID",
				zap.String("layer", logLayer),
				zap.String("function", logGetUserInstituteByID),
				zap.Int64("profileID", profileID),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetInstancesByProfileID: %w", err)
		}
		instances = append(instances, instanceTaken)
	}
	r.logger.Info("GetInstancesByProfileID Success",
		zap.String("layer", logLayer),
		zap.String("function", logGetUserInstituteByID),
		zap.Int64("profileID", profileID),
	)
	return instances, nil
}

func (r *UserCourseInstanceRepo) AddCourseInstance(ctx context.Context,
	userCourseInstance *usercourseinstance.UserCourseInstance) error {
	_, err := r.pool.Exec(ctx, queryAddCourseInstance, userCourseInstance.ProfileID, userCourseInstance.CourseInstanceID)
	if err != nil {
		r.logger.Error("AddCourseInstance",
			zap.String("layer", logLayer),
			zap.String("function", logAddCourseInstance),
			zap.Error(err),
		)
		return fmt.Errorf("AddCourseInstance: %w", err)
	}
	r.logger.Info("AddCourseInstance Success",
		zap.String("layer", logLayer),
		zap.String("function", logAddCourseInstance),
		zap.Int("profileID", userCourseInstance.ProfileID),
	)
	return nil
}
