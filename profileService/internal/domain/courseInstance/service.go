package courseInstance

import "context"

type Service interface {
	GetCourseInstanceByID(ctx context.Context, courseID int64) (*CourseInstance, error)
	AddCourseInstance(ctx context.Context, course *CourseInstance) error
	UpdateCourseInstanceByID(ctx context.Context, profile *CourseInstance) error
	GetInstancesByInstituteIDs(ctx context.Context, instituteIDs []int64) ([]int64, error)
	GetInstancesByAcademicYearIDs(ctx context.Context, academicYearIDs []int64) ([]int64, error)
	GetInstancesBySemesterIDs(ctx context.Context, semesterIDs []int64) ([]int64, error)
	GetInstancesByProgramIDs(ctx context.Context, programIDs []int64) ([]int64, error)
	GetInstancesByAllocationStatus(ctx context.Context, isAllocDone bool) ([]int64, error)
	GetInstancesByYear(ctx context.Context, year int64) ([]int64, error)
	GetInstancesByVersionID(ctx context.Context, versionID int64) ([]int64, error)
}
