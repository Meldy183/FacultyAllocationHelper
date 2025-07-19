package courseInstance

import "context"

type Repository interface {
	AddNewCourseInstance(ctx context.Context, profile *CourseInstance) error
	GetCourseInstanceByID(ctx context.Context, profileID int64) (*CourseInstance, error)
	UpdateCourseInstanceByID(ctx context.Context, profile *CourseInstance) error
	GetInstancesIDsByInstituteIDs(ctx context.Context, instituteIDs []int64) ([]int64, error)
	GetInstancesIDsByAcademicYearIDs(ctx context.Context, academicYearIDs []int64) ([]int64, error)
	GetInstancesIDsBySemesterIDs(ctx context.Context, semesterIDs []int64) ([]int64, error)
	GetInstancesIDsByProgramIDs(ctx context.Context, proframIDs []int64) ([]int64, error)
}
