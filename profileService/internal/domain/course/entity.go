package course

type Course struct {
	CourseID               int64
	Name                   string
	IsElective             *bool
	OfficialName           *string
	ResponsibleInstituteID int64
	LecHours               *int64
	LabHours               *int64
}
