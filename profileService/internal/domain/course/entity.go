package course

// "time"

type Course struct {
	CourseID     int64
	Name         string
	OfficialName *string
	LecHours     *int
	LabHours     *int
}

func NewCourse(
	id int64,
	name string,
	off_name *string,
	lec_hours, lab_hours *int,
) (*Course, error) {
	return &Course{
		CourseID:     id,
		Name:         name,
		OfficialName: off_name,
		LecHours:     lec_hours,
		LabHours:     lab_hours,
	}, nil
}
