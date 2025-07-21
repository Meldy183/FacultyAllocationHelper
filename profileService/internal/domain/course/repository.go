package course

import "context"

type Repository interface {
	AddNewCourse(ctx context.Context, profile *Course) error
	GetCourseByID(ctx context.Context, profileID int64) (*Course, error)
	UpdateCourseByID(ctx context.Context, profile *Course) error
}
