package course

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetCourseByID(ctx context.Context, courseID uuid.UUID) (*Course, error)
	AddCourse(ctx context.Context, course *Course) error
	UpdateCourseByID(ctx context.Context, id uuid.UUID, course *Course) error
}
