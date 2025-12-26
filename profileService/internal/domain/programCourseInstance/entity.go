package programcourseinstance

import "github.com/google/uuid"

type ProgramCourseInstance struct {
	ProgramCourseID  uuid.UUID
	ProgramID        uuid.UUID
	CourseInstanceID uuid.UUID
}
