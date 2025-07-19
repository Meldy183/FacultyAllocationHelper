package courseInstance

import "context"

type Repository interface {
	AddNewCourseInstance(ctx context.Context, profile *CourseInstance) error
	GetCourseInstanceByID(ctx context.Context, profileID int64) (*CourseInstance, error)
	UpdateCourseInstanceByID(ctx context.Context, profile *CourseInstance) error
	GetInstanceByInstituteID(ctx context.Context, instituteID int64) (*CourseInstance, error)
}
