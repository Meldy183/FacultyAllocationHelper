package courseInstance

import "context"

type Repository interface {
	AddNewCourseInstance(ctx context.Context, course *CourseInstance) (int64, error)
	GetCourseInstanceByID(ctx context.Context, courseID int64) (*CourseInstance, error)
	UpdateCourseInstanceByID(ctx context.Context, id int64, course *CourseInstance) error
	GetInstancesIDsByInstituteIDs(ctx context.Context, instituteIDs []int64) ([]int64, error)
	GetInstancesIDsByAcademicYearIDs(ctx context.Context, academicYearIDs []int64) ([]int64, error)
	GetInstancesIDsBySemesterIDs(ctx context.Context, semesterIDs []int64) ([]int64, error)
	GetInstancesIDsByProgramIDs(ctx context.Context, programIDs []int64) ([]int64, error)
	GetInstancesByAllocationStatus(ctx context.Context) ([]int64, error)
	GetInstancesByYear(ctx context.Context, year int64) ([]int64, error)
	GetInstancesByVersionID(ctx context.Context, versionID int64) ([]int64, error)
	GetAllInstancesIDs(ctx context.Context) ([]int64, error)
}
