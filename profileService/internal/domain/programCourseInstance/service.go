package programcourseinstance

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetProgramCourseInstancesByCourseID(ctx context.Context, instanceID uuid.UUID) ([]*ProgramCourseInstance, error)
	AddProgramToCourseInstance(ctx context.Context, programCourseInstance *ProgramCourseInstance) error
}
