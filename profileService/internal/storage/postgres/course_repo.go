package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/course"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ course.Repository = (*CourseRepo)(nil)

type CourseRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewCourseRepo(pool *pgxpool.Pool, logger *zap.Logger) *CourseRepo {
	return &CourseRepo{pool: pool, logger: logger}
}

const (
	queryGetCourseByID = `
		SELECT course_id, name, official_name, responsible_institute_id, lec_hours, lab_hours, is_elective
		FROM course
		WHERE course_id = $1
	`

	queryInsertCourse = `
		INSERT INTO course (
			name, official_name, responsible_institute_id, lec_hours, lab_hours, is_elective
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING course_id
	`

	queryUpdateCourseByID = `
		UPDATE course
		SET name = $1, official_name = $2,
		    responsible_institute_id = $3, lec_hours = $4, lab_hours = $5, is_elective = $6
		WHERE course_id = $7
	`
)

func (r *CourseRepo) GetCourseByID(ctx context.Context, courseID int64) (*course.Course, error) {
	row := r.pool.QueryRow(ctx, queryGetCourseByID, courseID)
	var course course.Course
	err := row.Scan(
		&course.CourseID,
		&course.Name,
		&course.OfficialName,
		&course.ResponsibleInstituteID,
		&course.LecHours,
		&course.LabHours,
		&course.IsElective,
	)
	if err != nil {
		r.logger.Error("Error getting course",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCourseByID),
			zap.Int64("courseID", courseID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetCourseByID failed: %w", err)
	}
	r.logger.Info("Course found",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetCourseByID),
		zap.Int64("courseID", courseID),
	)
	return &course, nil
}

func (r *CourseRepo) AddNewCourse(ctx context.Context, course *course.Course) error {
	err := r.pool.QueryRow(ctx, queryInsertCourse,
		course.Name,
		"",
		course.ResponsibleInstituteID,
		course.LecHours,
		course.LabHours,
		course.IsElective,
	).Scan(&course.CourseID)
	if err != nil {
		r.logger.Error("Error creating course",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogAddNewCourse),
			zap.Int64("courseID", course.CourseID),
			zap.Error(err),
		)
		return fmt.Errorf("AddNewCourse failed: %w", err)
	}
	r.logger.Info("Course profile created",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogAddNewCourse),
		zap.Int64("courseID", course.CourseID),
	)
	return nil
}

func (r *CourseRepo) UpdateCourseByID(ctx context.Context, id int64, course *course.Course) error {
	_, err := r.pool.Exec(ctx, queryUpdateCourseByID,
		course.Name,
		course.OfficialName,
		course.ResponsibleInstituteID,
		course.LecHours,
		course.LabHours,
		course.IsElective,
		id,
	)
	if err != nil {
		r.logger.Error("Error editing course",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogUpdateCourseByID),
			zap.Int64("courseID", course.CourseID),
			zap.Error(err),
		)
		return fmt.Errorf("AddNewCourse failed: %w", err)
	}
	r.logger.Info("Course profile updated",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogUpdateCourseByID),
		zap.Int64("courseID", course.CourseID),
	)
	return nil
}
