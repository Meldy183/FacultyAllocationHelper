package course

type Course struct {
	CourseID     int64
	Name         string
	OfficialName *string
	LecHours     *int
	LabHours     *int
}

func NewCourse(
	courseID int64,
	name string,
	offName *string,
	lecHours, labHours *int,
) (*Course, error) {
	return &Course{
		CourseID:     courseID,
		Name:         name,
		OfficialName: offName,
		LecHours:     lecHours,
		LabHours:     labHours,
	}, nil
}
