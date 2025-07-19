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

	queryGetInstancesByInstituteIDs = `
		SELECT ci.instance_id, ci.course_id, ci.semester_id, ci.year, ci.mode, ci.academic_year_id, ci.hardness_coefficient, ci.form, ci.groups_needed, ci.groups_taken, ci.pi_allocation_status, ci.ti_allocation_status
		FROM institute_course_link icl
		JOIN course c ON icl.course_id = c.course_id 
		LEFT JOIN course_instance ci ON ci.course_id = c.course_id
		WHERE icl.responsible_institute_id = ANY ($1)
		ORDER BY ci.instance_id
	`
	queryGetInstancesByAcademicYearIDs = `
		SELECT instance_id, course_id, semester_id, year, mode, academic_year_id, hardness_coefficient, form, groups_needed, groups_taken, pi_allocation_status, ti_allocation_status
		FROM course_instance
		WHERE academic_year_id = ANY ($1) 
		ORDER BY instance_id
	`
	queryGetInstancesBySemesterIDs = `
		SELECT instance_id, course_id, semester_id, year, mode, academic_year_id, hardness_coefficient, form, groups_needed, groups_taken, pi_allocation_status, ti_allocation_status
		FROM course_instance
		WHERE semester_id = ANY ($1) 
		ORDER BY instance_id
	`

	queryGetInstancesByProgramIDs = `
		SELECT 
			ci.instance_id, ci.course_id, ci.semester_id, ci.year, ci.mode, ci.academic_year_id, ci.hardness_coefficient, ci.form, ci.groups_needed, ci.groups_taken, ci.pi_allocation_status, ci.ti_allocation_status
		FROM program_course_instance pci JOIN course_instance ci ON pci.instance_id = ci.instance_id
		WHERE pci.program_id = $1
		ORDER BY ci.instance_id;
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

func (r *CourseInstanceRepo) GetInstancesByInstituteIDs(ctx context.Context, instituteIDs []int64) ([]*courseInstance.CourseInstance, error) {
	rows, err := r.pool.Query(ctx, queryGetInstancesByInstituteIDs, instituteIDs)
	if err != nil {
		r.logger.Error("Error getting courseInstances by instituteIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCourseInstanceByInstituteID),
			zap.String("instituteIDs", fmt.Sprintf("%v", instituteIDs)),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetInstancesByInstituteIDs failed: %w", err)
	}
	defer rows.Close()
	var instances []*courseInstance.CourseInstance
	for rows.Next() {
		var courseInstance courseInstance.CourseInstance
		err := rows.Scan(
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
			r.logger.Error("Error getting courseInstances by instituteIDs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetCourseInstanceByInstituteID),
				zap.String("instituteIDs", fmt.Sprintf("%v", instituteIDs)),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetInstancesByInstituteIDs failed: %w", err)
		}
		instances = append(instances, &courseInstance)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error getting courseInstance by instituteID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCourseInstanceByInstituteID),
			zap.String("instituteIDs", fmt.Sprintf("%v", instituteIDs)),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetCourseInstanceByInstituteID failed: %w", err)
	}
	r.logger.Info("CourseInstances found by instituteIDs",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetCourseInstanceByInstituteID),
		zap.Int("instancesLen", len(instances)),
	)
	return instances, nil
}

func (r *CourseInstanceRepo) GetInstancesByAcademicYearIDs(ctx context.Context, academicYearIDs []int64) ([]*courseInstance.CourseInstance, error) {
	rows, err := r.pool.Query(ctx, queryGetInstancesByAcademicYearIDs, academicYearIDs)
	if err != nil {
		r.logger.Error("Error getting courseInstances by academicYearIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCourseInstanceByacademicYearID),
			zap.String("academicYearIDs", fmt.Sprintf("%v", academicYearIDs)),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetInstancesByAcademicYearIDs failed: %w", err)
	}
	defer rows.Close()
	var instances []*courseInstance.CourseInstance
	for rows.Next() {
		var courseInstance courseInstance.CourseInstance
		err := rows.Scan(
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
			r.logger.Error("Error getting courseInstances by academicYearIDs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetCourseInstanceByacademicYearID),
				zap.String("academicYearIDs", fmt.Sprintf("%v", academicYearIDs)),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetInstancesByAcademicYearIDs failed: %w", err)
		}
		instances = append(instances, &courseInstance)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error getting courseInstance by academicYearIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCourseInstanceByacademicYearID),
			zap.String("academicYearIDs", fmt.Sprintf("%v", academicYearIDs)),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetCourseInstanceByAcademicYearIDs failed: %w", err)
	}
	r.logger.Info("CourseInstances found by academicYearIDs",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetCourseInstanceByacademicYearID),
		zap.Int("instancesLen", len(instances)),
	)
	return instances, nil
}

func (r *CourseInstanceRepo) GetInstancesBySemesterIDs(ctx context.Context, semesterIDs []int64) ([]*courseInstance.CourseInstance, error) {
	rows, err := r.pool.Query(ctx, queryGetInstancesBySemesterIDs, semesterIDs)
	if err != nil {
		r.logger.Error("Error getting courseInstances by semesterIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetInstanceBySemesterID),
			zap.String("semesterIDs", fmt.Sprintf("%v", semesterIDs)),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetInstancesBySemesterIDs failed: %w", err)
	}
	defer rows.Close()
	var instances []*courseInstance.CourseInstance
	for rows.Next() {
		var courseInstance courseInstance.CourseInstance
		err := rows.Scan(
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
			r.logger.Error("Error getting courseInstances by semesterIDs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetInstanceBySemesterID),
				zap.String("semesterIDs", fmt.Sprintf("%v", semesterIDs)),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetInstancesBySemesterIDs failed: %w", err)
		}
		instances = append(instances, &courseInstance)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error getting courseInstance by semesterIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetInstanceBySemesterID),
			zap.String("semesterIDs", fmt.Sprintf("%v", semesterIDs)),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetCourseInstanceBySemesterIDs failed: %w", err)
	}
	r.logger.Info("CourseInstances found by semesterIDs",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetInstanceBySemesterID),
		zap.Int("instancesLen", len(instances)),
	)
	return instances, nil
}

func (r *CourseInstanceRepo) GetInstancesByProgramIDs(ctx context.Context, programIDs []int64) ([]*courseInstance.CourseInstance, error) {
	rows, err := r.pool.Query(ctx, queryGetInstancesByProgramIDs, programIDs)
	if err != nil {
		r.logger.Error("Error getting courseInstances by programIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCourseInstanceByProgramID),
			zap.String("programIDs", fmt.Sprintf("%v", programIDs)),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetInstancesByProgramIDs failed: %w", err)
	}
	defer rows.Close()
	var instances []*courseInstance.CourseInstance
	for rows.Next() {
		var courseInstance courseInstance.CourseInstance
		err := rows.Scan(
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
			r.logger.Error("Error getting courseInstances by programIDs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetCourseInstanceByProgramID),
				zap.String("programIDs", fmt.Sprintf("%v", programIDs)),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetInstancesByProgramIDs failed: %w", err)
		}
		instances = append(instances, &courseInstance)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error getting courseInstance by programIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCourseInstanceByProgramID),
			zap.String("programIDs", fmt.Sprintf("%v", programIDs)),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetCourseInstanceByProgramIDs failed: %w", err)
	}
	r.logger.Info("CourseInstances found by programIDs",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetCourseInstanceByProgramID),
		zap.Int("instancesLen", len(instances)),
	)
	return instances, nil
}
