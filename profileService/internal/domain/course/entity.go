package course

type Course struct {
	CourseID               int64
	Name                   string
	OfficialName           *string
	ResponsibleInstituteID int64
	LecHours               *int
	LabHours               *int
}

func NewCourse(
	courseID, responsibleInstituteID int64,
	name string,
	offName *string,
	lecHours, labHours *int,
) (*Course, error) {
	return &Course{
		CourseID:               courseID,
		Name:                   name,
		OfficialName:           offName,
		ResponsibleInstituteID: responsibleInstituteID,
		LecHours:               lecHours,
		LabHours:               labHours,
	}, nil
}
