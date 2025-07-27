package course

import "context"

type Repository interface {
	AddNewCourse(ctx context.Context, course *Course) error
	AddNewCourseIDReturned(ctx context.Context, course *Course) (int64, error)
	GetCourseByID(ctx context.Context, courseID int64) (*Course, error)
	UpdateCourseByID(ctx context.Context, id int64, course *Course) error
}
