package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/courseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ courseInstance.Repository = (*CourseInstanceRepo)(nil)

type CourseInstanceRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewCourseInstanceRepo(pool *pgxpool.Pool, logger *zap.Logger) *CourseInstanceRepo {
	return &CourseInstanceRepo{pool: pool, logger: logger}
}

const (
	queryGetCourseInstanceByID = `
		SELECT instance_id, course_id, semester, year, mode, academic_year_id, hardness_coefficient, form, groups_needed, groups_taken, pi_allocation_status, ti_allocation_status
		FROM course_instance
		WHERE instance_id = $1
	`

	queryInsertCourseInstance = `
		INSERT INTO course_instance (
			semester_id, year, mode, academic_year_id, hardness_coefficient, form, groups_needed, groups_taken, pi_allocation_status, ti_allocation_status
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING instance_id
	`

	queryUpdateCourseInstanceByID = `
		UPDATE course_instance
		SET brief_name = $1, official_name = $2,
		    responsible_institute_id = $3, lec_hours = $4, lab_hours = $5
		WHERE course_instance_id = $6
	`

	queryGetInstanceByInstituteID = `
		SELECT instance_id, course_id, semester_id, year, mode, academic_year_id, hardness_coefficient, form, groups_needed, groups_taken, pi_allocation_status, ti_allocation_status
		FROM institute_course_link icl JOIN course_instance ci ON icl.course_instance_id = ci.course_instance_id
		WHERE icl.responsible_institute_id = $1
	`
	queryGetInstanceByAcademicYearID = `
		SELECT instance_id, course_id, semester_id, year, mode, academic_year_id, hardness_coefficient, form, groups_needed, groups_taken, pi_allocation_status, ti_allocation_status
		FROM course_instance
		WHERE academic_year_id = $1
	`
	queryGetInstanceBySemesterID = `
		SELECT instance_id, course_id, semester_id, year, mode, academic_year_id, hardness_coefficient, form, groups_needed, groups_taken, pi_allocation_status, ti_allocation_status
		FROM course_instance
		WHERE semester_id = $1
	`
)

func (r *CourseInstanceRepo) GetCourseInstanceByID(ctx context.Context, courseInstanceID int64) (*courseInstance.CourseInstance, error) {
	row := r.pool.QueryRow(ctx, queryGetCourseInstanceByID, courseInstanceID)
	var courseInstance courseInstance.CourseInstance
	err := row.Scan(
		&courseInstance.InstanceID,
		&courseInstance.CourseID,
		&courseInstance.SemesterID,
		&courseInstance.Year,
		&courseInstance.Mode,
		&courseInstance.AcademicYearID,
		&courseInstance.HardnessCoefficient,
		&courseInstance.Form,
		&courseInstance.GroupsNeeded,
		&courseInstance.GroupsTaken,
		&courseInstance.PIAllocationStatus,
		&courseInstance.TIAllocationStatus,
	)
	if err != nil {
		r.logger.Error("Error getting courseInstance",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCourseInstanceByID),
			zap.Int64("courseInstanceID", courseInstanceID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetCourseInstanceByID failed: %w", err)
	}
	r.logger.Info("CourseInstance found",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetCourseInstanceByID),
		zap.Int64("courseInstanceID", courseInstanceID),
	)
	return &courseInstance, nil
}

func (r *CourseInstanceRepo) AddNewCourseInstance(ctx context.Context, courseInstance *courseInstance.CourseInstance) error {
	err := r.pool.QueryRow(ctx, queryInsertCourseInstance,
		courseInstance.SemesterID,
		courseInstance.Year,
		courseInstance.Mode,
		courseInstance.AcademicYearID,
		courseInstance.HardnessCoefficient,
		courseInstance.Form,
		courseInstance.GroupsNeeded,
		courseInstance.GroupsTaken,
		courseInstance.PIAllocationStatus,
		courseInstance.TIAllocationStatus,
	).Scan(&courseInstance.InstanceID)
	if err != nil {
		r.logger.Error("Error creating courseInstance",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogAddNewCourseInstance),
			zap.Int64("courseInstanceeID", courseInstance.InstanceID),
			zap.Error(err),
		)
		return fmt.Errorf("AddNewCourseInstance failed: %w", err)
	}
	r.logger.Info("CourseInstance created",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogAddNewCourseInstance),
		zap.Int64("courseInstanceID", courseInstance.InstanceID),
	)
	return nil
}

func (r *CourseInstanceRepo) UpdateCourseInstanceByID(ctx context.Context, courseInstance *courseInstance.CourseInstance) error {
	_, err := r.pool.Exec(ctx, queryUpdateCourseInstanceByID,
		courseInstance.SemesterID,
		courseInstance.Year,
		courseInstance.Mode,
		courseInstance.AcademicYearID,
		courseInstance.HardnessCoefficient,
		courseInstance.Form,
		courseInstance.GroupsNeeded,
		courseInstance.GroupsTaken,
		courseInstance.PIAllocationStatus,
		courseInstance.TIAllocationStatus,
	)
	if err != nil {
		r.logger.Error("Error editing courseInstance",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogUpdateCourseInstanceByID),
			zap.Int64("courseInstanceID", courseInstance.InstanceID),
			zap.Error(err),
		)
		return fmt.Errorf("UpdateCourseInstance failed: %w", err)
	}
	r.logger.Info("CourseInstance updated",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogUpdateCourseInstanceByID),
		zap.Int64("courseInstanceID", courseInstance.InstanceID),
	)
	return nil
}

func (r *CourseInstanceRepo) GetInstancesByInstituteIDs(ctx context.Context, instituteID int64) (*courseInstance.CourseInstance, error) {
	row := r.pool.QueryRow(ctx, queryGetInstanceByInstituteID, instituteID)
	var courseInstance courseInstance.CourseInstance
	err := row.Scan(
		&courseInstance.InstanceID,
		&courseInstance.CourseID,
		&courseInstance.SemesterID,
		&courseInstance.Year,
		&courseInstance.Mode,
		&courseInstance.AcademicYearID,
		&courseInstance.HardnessCoefficient,
		&courseInstance.Form,
		&courseInstance.GroupsNeeded,
		&courseInstance.GroupsTaken,
		&courseInstance.PIAllocationStatus,
		&courseInstance.TIAllocationStatus,
	)
	if err != nil {
		r.logger.Error("Error getting courseInstance by instituteID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCourseInstanceByInstituteID),
			zap.Int64("instituteID", instituteID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetCourseInstanceByInstituteID failed: %w", err)
	}
	r.logger.Info("CourseInstance found by instituteID",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetCourseInstanceByInstituteID),
		zap.Int64("instituteID", instituteID),
	)
	return &courseInstance, nil
}

func (r *CourseInstanceRepo) GetInstancesByAcademicYearIDs(ctx context.Context, academicYearID int64) (*courseInstance.CourseInstance, error) {
	row := r.pool.QueryRow(ctx, queryGetInstanceByAcademicYearID, academicYearID)
	var courseInstance courseInstance.CourseInstance
	err := row.Scan(
		&courseInstance.InstanceID,
		&courseInstance.CourseID,
		&courseInstance.SemesterID,
		&courseInstance.Year,
		&courseInstance.Mode,
		&courseInstance.AcademicYearID,
		&courseInstance.HardnessCoefficient,
		&courseInstance.Form,
		&courseInstance.GroupsNeeded,
		&courseInstance.GroupsTaken,
		&courseInstance.PIAllocationStatus,
		&courseInstance.TIAllocationStatus,
	)
	if err != nil {
		r.logger.Error("Error getting courseInstance by instituteID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCourseInstanceByacademicYearID),
			zap.Int64("academicYearID", academicYearID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetCourseInstanceByInstituteID failed: %w", err)
	}
	r.logger.Info("CourseInstance found by instituteID",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetCourseInstanceByacademicYearID),
		zap.Int64("academicYearID", academicYearID),
	)
	return &courseInstance, nil
}

func (r *CourseInstanceRepo) GetInstancesBySemesterIDs(ctx context.Context, semesterID int64) (*courseInstance.CourseInstance, error) {
	row := r.pool.QueryRow(ctx, queryGetInstanceBySemesterID, semesterID)
	var courseInstance courseInstance.CourseInstance
	err := row.Scan(
		&courseInstance.InstanceID,
		&courseInstance.CourseID,
		&courseInstance.SemesterID,
		&courseInstance.Year,
		&courseInstance.Mode,
		&courseInstance.AcademicYearID,
		&courseInstance.HardnessCoefficient,
		&courseInstance.Form,
		&courseInstance.GroupsNeeded,
		&courseInstance.GroupsTaken,
		&courseInstance.PIAllocationStatus,
		&courseInstance.TIAllocationStatus,
	)
	if err != nil {
		r.logger.Error("Error getting courseInstance by semesterID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetInstanceBySemesterID),
			zap.Int64("semesterID", semesterID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetCourseInstanceByInstituteID failed: %w", err)
	}
	r.logger.Info("CourseInstance found by semesterID",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetInstanceBySemesterID),
		zap.Int64("semesterID", semesterID),
	)
	return &courseInstance, nil
}
