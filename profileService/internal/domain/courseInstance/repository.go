package courseInstance

import "context"

type Repository interface {
	AddNewCourseInstance(ctx context.Context, profile *CourseInstance) error
	GetCourseInstanceByID(ctx context.Context, profileID int64) (*CourseInstance, error)
	UpdateCourseInstanceByID(ctx context.Context, profile *CourseInstance) error
	GetInstancesByInstituteIDs(ctx context.Context, instituteIDs []int64) ([]int64, error)
	GetInstancesByAcademicYearIDs(ctx context.Context, academicYearIDs []int64) ([]int64, error)
	GetInstancesBySemesterIDs(ctx context.Context, semesterIDs []int64) ([]int64, error)
	GetInstancesByProgramIDs(ctx context.Context, proframIDs []int64) ([]int64, error)
}
