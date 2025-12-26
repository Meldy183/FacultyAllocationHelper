package courseInstitute

import "github.com/google/uuid"

type InstituteCourseLink struct {
	CourseInstituteID uuid.UUID
	InstituteID       uuid.UUID
	CourseID          uuid.UUID
}
