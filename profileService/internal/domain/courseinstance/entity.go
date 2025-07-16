package courseinstance

import UserProfileVersion


type CourseInstance struct {
	InstanceID     int64
	Year           int
	CourseID       int64
	SemesterID     int
	AcademicYearID int
	Form           *Form
	Mode           *Mode
	GroupsNeeded   int
	GroupsTaken    *int
	PIAllocationStatus Status
}

type Status string

const (
	ID Status = 
	StatusNA Status = "Not Allocated"
	StatusNN   Status = "Not needed"
)

type Form string

const (
	FormBlock1 Form = "First Block"
	FormBlock2 Form = "Second Block"
	FormFull   Form = "Full"
)

type Mode string

const (
	ModeOnsite Mode = "Onsite"
	ModeMixed  Mode = "Mixed"
	ModeRemote Mode = "Remote"
)

func NewCourse(
	id int64,
	name string,
	off_name *string,
	lec_hours, lab_hours *int,
) (*CourseInstance, error) {
	return &CourseInstance{
		CourseID:     id,
		Name:         name,
		OfficialName: off_name,
		LecHours:     lec_hours,
		LabHours:     lab_hours,
	}, nil
}
