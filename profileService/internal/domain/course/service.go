package course

import "context"

type Service interface {
	GetCourseByID(ctx context.Context, courseID int64) (*Course, error)
	AddCourse(ctx context.Context, course *Course) error
}
