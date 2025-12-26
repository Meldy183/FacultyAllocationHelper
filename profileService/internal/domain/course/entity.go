package course

import "github.com/google/uuid"

type Course struct {
	CourseID               uuid.UUID
	Name                   string
	IsElective             *bool
	OfficialName           *string
	ResponsibleInstituteID uuid.UUID
	LecHours               *int64
	LabHours               *int64
}
