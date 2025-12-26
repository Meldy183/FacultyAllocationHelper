package CompleteCourse

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetFullCourseInfoByID(ctx context.Context, instanceID uuid.UUID) (*FullCourse, error)
	AddFullCourse(ctx context.Context, fullCourse *FullCourse) error
}
