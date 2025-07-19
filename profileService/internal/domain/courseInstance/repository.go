package courseInstance

import "context"

type Repository interface {
	AddNewCourseInstance(ctx context.Context, profile *CourseInstance) error
	GetCourseInstanceByID(ctx context.Context, profileID int64) (*CourseInstance, error)
	UpdateCourseInstanceByID(ctx context.Context, profile *CourseInstance) error
	GetInstancesByInstituteIDs(ctx context.Context, instituteIDs []int64) ([]*CourseInstance, error)
	GetInstancesByAcademicYearIDs(ctx context.Context, academicYearIDs []int64) ([]*CourseInstance, error)
	GetInstancesBySemesterIDs(ctx context.Context, semesterIDs []int64) ([]*CourseInstance, error)
}
