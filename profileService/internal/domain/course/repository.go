package course

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	AddNewCourse(ctx context.Context, course *Course) error
	GetCourseByID(ctx context.Context, courseID uuid.UUID) (*Course, error)
	UpdateCourseByID(ctx context.Context, id uuid.UUID, course *Course) error
}
