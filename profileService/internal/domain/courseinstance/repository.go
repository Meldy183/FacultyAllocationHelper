package courseinstance

import "context"

type Repository interface {
	AddNewCourse(ctx context.Context, profile *CourseInstance) error
	GetCourseByID(ctx context.Context, profileID int64) (*CourseInstance, error)
	UpdateCourseByID(ctx context.Context, profile *CourseInstance) error
}
