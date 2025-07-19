package course

type Course struct {
	CourseID               int64
	Name                   string
	OfficialName           *string
	ResponsibleInstituteID int64
	LecHours               *int
	LabHours               *int
}
