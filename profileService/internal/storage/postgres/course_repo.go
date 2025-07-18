package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/course"
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
// queryGetCourseByID = `
// 	SELECT course_id, brief_name, official_name, academic_year_name, semester_name
// 	FROM course
// 	WHERE course_id = $1
// `

// 	queryInsertUserProfile = `
// 		INSERT INTO user_profile (
// 			email, english_name, alias
// 		)
// 		VALUES ($1, $2, $3)
// 		RETURNING profile_id
// 	`

//	queryUpdateUserProfile = `
//		UPDATE user_profile
//		SET email = $1, english_name = $2,
//		    russian_name = $3, alias = $4, start_date = $5, end_date = $6
//		WHERE profile_id = $7
//	`
//	queryGetProfileIDsByInstituteIDs = `SELECT profile_id FROM user_institute WHERE institute_id = ANY($1)
//
// ORDER BY profile_id`
//
//	queryGetProfileIDsByPositionIDs = `SELECT profile_id from user_profile_version where position_id = ANY($1)
//
// ORDER BY profile_id`
)

func (r *CourseRepo) GetCourseByID(ctx context.Context, profileID int64) (*course.Course, error) {

}

func (r *CourseRepo) UpdateCourseByID(ctx context.Context, course *course.Course) error {

}

func (r *CourseRepo) AddNewCourse(ctx context.Context, course *course.Course) error {

}
