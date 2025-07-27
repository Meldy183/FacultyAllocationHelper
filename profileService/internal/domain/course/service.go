package course

import "context"

type Service interface {
	GetCourseByID(ctx context.Context, courseID int64) (*Course, error)
	AddCourse(ctx context.Context, course *Course) error
	AddNewCourseIDReturned(ctx context.Context, course *Course) (int64, error)
	UpdateCourseByID(ctx context.Context, id int64, course *Course) error
}
