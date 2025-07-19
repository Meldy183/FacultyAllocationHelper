package courseInstance

type CourseInstance struct {
	InstanceID          int64
	Year                int
	CourseID            int64
	SemesterID          int
	AcademicYearID      int
	HardnessCoefficient *float64
	Form                *Form
	Mode                *Mode
	GroupsNeeded        int
	GroupsTaken         *int
	PIAllocationStatus  Status
	TIAllocationStatus  Status
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
