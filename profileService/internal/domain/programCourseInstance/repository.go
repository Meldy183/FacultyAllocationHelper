package programcourseinstance

import "context"

type Repository interface {
	GetProgramCourseInstancesByCourseID(ctx context.Context, courseID int64) ([]*ProgramCourseInstance, error)
}
