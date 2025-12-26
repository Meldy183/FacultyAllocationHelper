package programcourseinstance

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetProgramCourseInstancesByCourseID(ctx context.Context, courseID uuid.UUID) ([]*ProgramCourseInstance, error)
	AddProgramToCourseInstance(ctx context.Context, programCourseInstance *ProgramCourseInstance) error
}
