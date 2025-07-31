package courseInstance

type CourseInstance struct {
	InstanceID          int64
	Year                int64
	CourseID            int64
	SemesterID          int64
	AcademicYearID      int64
	HardnessCoefficient *float64
	Form                *Form
	Mode                *Mode
	GroupsNeeded        int64
	GroupsTaken         *int64
	PIAllocationStatus  *Status
	TIAllocationStatus  *Status
}

type Form string

const (
	FormBlock1 Form = "First Block"
	FormBlock2 Form = "Second Block"
	FormFull   Form = "Full"
)

type Status string

const (
	StatusAllocated    Status = "Allocated"
	StatusNotAllocated Status = "Not allocated"
	StatusNotNeeded    Status = "Not needed"
)

type Mode string

const (
	ModeOnsite Mode = "Onsite"
	ModeMixed  Mode = "Mixed"
	ModeRemote Mode = "Remote"
)

func NewStatusDefault() *Status {
	x := StatusNotAllocated
	return &x
}

func NewModeDefault() *Mode {
	x := ModeOnsite
	return &x
}
func NewFormDefault() *Form {
	x := FormFull
	return &x
}
