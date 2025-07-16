package courseinstance

type CourseInstance struct {
	InstanceID         int64
	Year               int
	CourseID           int64
	SemesterID         int
	AcademicYearID     int
	Form               *Form
	Mode               *Mode
	GroupsNeeded       int
	GroupsTaken        *int
	PIAllocationStatus Status
	TIAllocationStatus Status
}

type StatusState int

const (
	StatusNotAllocated StatusState = iota
	StatusNotNeeded
	StatusHasValue
)

type Status struct {
	State StatusState
	Value int
}

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
	instance_id, course_id int64,
	year, semester_id, academic_year_id, groups_needed int,
	groups_taken *int,
	form *Form,
	mode *Mode,
	pi_alloc_Stat Status,
	ti_alloc_status Status,

) (*CourseInstance, error) {
	return &CourseInstance{
		InstanceID:         instance_id,
		Year:               year,
		CourseID:           course_id,
		SemesterID:         semester_id,
		AcademicYearID:     academic_year_id,
		Form:               form,
		Mode:               mode,
		GroupsNeeded:       groups_needed,
		GroupsTaken:        groups_taken,
		PIAllocationStatus: pi_alloc_Stat,
		TIAllocationStatus: ti_alloc_status,
	}, nil
}
