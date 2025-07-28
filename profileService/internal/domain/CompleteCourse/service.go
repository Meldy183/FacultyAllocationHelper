package CompleteCourse

import "context"

type Service interface {
	GetFullCourseInfoByID(ctx context.Context, instanceID int64) (*FullCourse, error)
	AddFullCourse(ctx context.Context, fullCourse *FullCourse) error
}
