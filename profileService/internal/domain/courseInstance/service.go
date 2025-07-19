package courseInstance

import "context"

type Service interface {
	GetCourseInstanceByID(ctx context.Context, courseID int64) (*CourseInstance, error)
	AddCourseInstance(ctx context.Context, course *CourseInstance) error
	UpdateCourseInstanceByID(ctx context.Context, profile *CourseInstance) error
	GetInstancesByInstitutes(ctx context.Context, instituteIDs []int64) ([]int64, error)
	GetInstancesByAcademicYears(ctx context.Context, academicYearIDs []int64) ([]int64, error)
	GetInstancesBySemesters(ctx context.Context, semesterIDs []int64) ([]int64, error)
	GetInstancesByPrograms(ctx context.Context, proframIDs []int64) ([]int64, error)
}
