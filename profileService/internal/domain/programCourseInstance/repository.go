package programcourseinstance

import "context"

type Repository interface {
	GetProgramCourseInstancesByID(ctx context.Context, code string) (*ProgramCourseInstance, error)
}
