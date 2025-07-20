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
		SELECT instance_id, course_id, semester_id, year, mode, academic_year_id, hardness_coefficient, form, groups_needed, groups_taken, pi_allocation_status, ti_allocation_status
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
		SELECT ci.instance_id
		FROM institute_course_link icl
		JOIN course c ON icl.course_id = c.course_id 
		LEFT JOIN course_instance ci ON ci.course_id = c.course_id
		WHERE icl.responsible_institute_id = ANY ($1)
		ORDER BY ci.instance_id
	`
	queryGetInstancesByAcademicYearIDs = `
		SELECT instance_id
		FROM course_instance
		WHERE academic_year_id = ANY ($1) 
		ORDER BY instance_id
	`
	queryGetInstancesBySemesterIDs = `
		SELECT instance_id
		FROM course_instance
		WHERE semester_id = ANY ($1) 
		ORDER BY instance_id
	`

	queryGetInstancesByProgramIDs = `
		SELECT ci.instance_id
		FROM program_course_instance pci JOIN course_instance ci ON pci.instance_id = ci.instance_id
		WHERE pci.program_id = $1
		ORDER BY ci.instance_id;
	`

	queryGetInstancesByAllocationStatus = `
		SELECT instance_id
		FROM course_instance
		WHERE groups_needed <> groups_taken
		ORDER BY instance_id
	`

	queryGetInstancesByYear = `
		SELECT instance_id
		FROM course_instance
		WHERE year = $1 
		ORDER BY instance_id
	`

	queryGetInstancesByVersionID = `
		SELECT ci.instance_id
		FROM course_instance ci JOIN staff s ON ci.instance_id = s.instance_id
		LEFT JOIN user_profile_version pv ON pv.profile_version_id = s.profile_version_id
		WHERE ci.profile_version_id = $1 
		ORDER BY instance_id
	`
	queryGetAllInstancesIDs = `
		SELECT instance_id
		FROM course_instance
		ORDER BY instance_id
	`
)

func (r *CourseInstanceRepo) GetAllInstancesIDs(ctx context.Context) ([]int64, error) {
	rows, err := r.pool.Query(ctx, queryGetAllInstancesIDs)
	if err != nil {
		r.logger.Error("Error getting all courseInstances IDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetAllInstancesIDs),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetAllInstancesIDs failed: %w", err)
	}
	defer rows.Close()
	var instancesIDs []int64
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			r.logger.Error("Error getting all courseInstances",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllInstancesIDs),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetAllInstancesIDs failed: %w", err)
		}
		instancesIDs = append(instancesIDs, id)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error getting all courseInstances",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetAllInstancesIDs),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetAllInstancesIDs failed: %w", err)
	}
	r.logger.Info("all CourseInstances found",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetAllInstancesIDs),
		zap.Int("instancesLen", len(instancesIDs)),
	)
	return instancesIDs, nil
}

func (r *CourseInstanceRepo) GetCourseInstanceByID(ctx context.Context, courseInstanceID int64) (*courseInstance.CourseInstance, error) {
	row := r.pool.QueryRow(ctx, queryGetCourseInstanceByID, courseInstanceID)
	var courseInstanceObj courseInstance.CourseInstance
	err := row.Scan(
		&courseInstanceObj.InstanceID,
		&courseInstanceObj.CourseID,
		&courseInstanceObj.SemesterID,
		&courseInstanceObj.Year,
		&courseInstanceObj.Mode,
		&courseInstanceObj.AcademicYearID,
		&courseInstanceObj.HardnessCoefficient,
		&courseInstanceObj.Form,
		&courseInstanceObj.GroupsNeeded,
		&courseInstanceObj.GroupsTaken,
		&courseInstanceObj.PIAllocationStatus,
		&courseInstanceObj.TIAllocationStatus,
	)
	if err != nil {
		r.logger.Error("Error getting courseInstanceObj",
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
	return &courseInstanceObj, nil
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

func (r *CourseInstanceRepo) GetInstancesIDsByInstituteIDs(ctx context.Context, instituteIDs []int64) ([]int64, error) {
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
	var instancesIDs []int64
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			r.logger.Error("Error getting courseInstances by instituteIDs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetCourseInstanceByInstituteID),
				zap.String("instituteIDs", fmt.Sprintf("%v", instituteIDs)),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetInstancesByInstituteIDs failed: %w", err)
		}
		instancesIDs = append(instancesIDs, id)
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
		zap.Int("instancesLen", len(instancesIDs)),
	)
	return instancesIDs, nil
}

func (r *CourseInstanceRepo) GetInstancesIDsByAcademicYearIDs(ctx context.Context, academicYearIDs []int64) ([]int64, error) {
	rows, err := r.pool.Query(ctx, queryGetInstancesByAcademicYearIDs, academicYearIDs)
	if err != nil {
		r.logger.Error("Error getting courseInstances by academicYearIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCourseInstanceByAcademicYearID),
			zap.String("academicYearIDs", fmt.Sprintf("%v", academicYearIDs)),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetInstancesByAcademicYearIDs failed: %w", err)
	}
	defer rows.Close()
	var instancesIDs []int64
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			r.logger.Error("Error getting courseInstances by academicYearIDs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetCourseInstanceByAcademicYearID),
				zap.String("academicYearIDs", fmt.Sprintf("%v", academicYearIDs)),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetInstancesByAcademicYearIDs failed: %w", err)
		}
		instancesIDs = append(instancesIDs, id)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error getting courseInstance by academicYearIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCourseInstanceByAcademicYearID),
			zap.String("academicYearIDs", fmt.Sprintf("%v", academicYearIDs)),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetCourseInstanceByAcademicYearIDs failed: %w", err)
	}
	r.logger.Info("CourseInstances found by academicYearIDs",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetCourseInstanceByAcademicYearID),
		zap.Int("instancesLen", len(instancesIDs)),
	)
	return instancesIDs, nil
}

func (r *CourseInstanceRepo) GetInstancesIDsBySemesterIDs(ctx context.Context, semesterIDs []int64) ([]int64, error) {
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
	var instancesIDs []int64
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			r.logger.Error("Error getting courseInstances by semesterIDs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetInstanceBySemesterID),
				zap.String("semesterIDs", fmt.Sprintf("%v", semesterIDs)),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetInstancesBySemesterIDs failed: %w", err)
		}
		instancesIDs = append(instancesIDs, id)
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
		zap.Int("instancesLen", len(instancesIDs)),
	)
	return instancesIDs, nil
}

func (r *CourseInstanceRepo) GetInstancesIDsByProgramIDs(ctx context.Context, programIDs []int64) ([]int64, error) {
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
	var instancesIDs []int64
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			r.logger.Error("Error getting courseInstances by programIDs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetCourseInstanceByProgramID),
				zap.String("programIDs", fmt.Sprintf("%v", programIDs)),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetInstancesByProgramIDs failed: %w", err)
		}
		instancesIDs = append(instancesIDs, id)
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
		zap.Int("instancesLen", len(instancesIDs)),
	)
	return instancesIDs, nil
}

func (r *CourseInstanceRepo) GetInstancesByAllocationStatus(ctx context.Context) ([]int64, error) {
	rows, err := r.pool.Query(ctx, queryGetInstancesByAllocationStatus)
	if err != nil {
		r.logger.Error("Error getting courseInstances by allocation status",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetInstancesIDsByAllocationStatus),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetInstancesByAllocationStatus failed: %w", err)
	}
	defer rows.Close()
	var instancesIDs []int64
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			r.logger.Error("Error getting courseInstances by allocation status",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetInstancesIDsByAllocationStatus),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetInstancesByAllocationStatus failed: %w", err)
		}
		instancesIDs = append(instancesIDs, id)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error getting courseInstance by allocation status",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetInstancesIDsByAllocationStatus),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetInstancesByAllocationStatus failed: %w", err)
	}
	r.logger.Info("CourseInstances found by allocation status",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetInstancesIDsByAllocationStatus),
		zap.Int("instancesLen", len(instancesIDs)),
	)
	return instancesIDs, nil
}

func (r *CourseInstanceRepo) GetInstancesByYear(ctx context.Context, year int) ([]int64, error) {
	rows, err := r.pool.Query(ctx, queryGetInstancesByYear, year)
	if err != nil {
		r.logger.Error("Error getting courseInstances by year",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetInstancesIDsByYear),
			zap.Int("year", year),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetInstancesIDsByYear failed: %w", err)
	}
	defer rows.Close()
	var instancesIDs []int64
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			r.logger.Error("Error getting courseInstances by year",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetInstancesIDsByYear),
				zap.Int("year", year),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetInstancesIDsByYear failed: %w", err)
		}
		instancesIDs = append(instancesIDs, id)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error getting courseInstance by year",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetInstancesIDsByYear),
			zap.Int("year", year),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetInstancesIDsByYear failed: %w", err)
	}
	r.logger.Info("CourseInstances found by year",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetInstancesIDsByYear),
		zap.Int("instancesLen", len(instancesIDs)),
	)
	return instancesIDs, nil
}

func (r *CourseInstanceRepo) GetInstancesByVersionID(ctx context.Context, versionID int64) ([]int64, error) {
	rows, err := r.pool.Query(ctx, queryGetInstancesByVersionID, versionID)
	if err != nil {
		r.logger.Error("Error getting courseInstances by versionID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCourseInstanceByVersionID),
			zap.Int64("versionID", versionID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetInstancesIDsByYear failed: %w", err)
	}
	defer rows.Close()
	var instancesIDs []int64
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			r.logger.Error("Error getting courseInstances by versionID",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetCourseInstanceByVersionID),
				zap.Int64("versionID", versionID),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetInstancesIDsByYear failed: %w", err)
		}
		instancesIDs = append(instancesIDs, id)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error getting courseInstance by versionID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCourseInstanceByVersionID),
			zap.Int64("versionID", versionID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetInstancesIDsByYear failed: %w", err)
	}
	r.logger.Info("CourseInstances found by versionID",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetCourseInstanceByVersionID),
		zap.Int("instancesLen", len(instancesIDs)),
	)
	return instancesIDs, nil
}
