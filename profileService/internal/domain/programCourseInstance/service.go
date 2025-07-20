package programcourseinstance

import "context"

type Service interface {
	GetProgramCourseInstancesByCourseID(ctx context.Context, instanceID int64) ([]*ProgramCourseInstance, error)
	AddProgramToCourseInstance(ctx context.Context, programCourseInstance *ProgramCourseInstance) error
}
