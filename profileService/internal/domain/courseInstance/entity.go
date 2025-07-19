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

func NewCourseInstance(
	instanceId, courseId int64,
	year, semesterId, academicYearId, groupsNeeded int,
	hardnessCoefficient *float64,
	groupsTaken *int,
	form *Form,
	mode *Mode,
	piAllocStat Status,
	tiAllocStatus Status,

) (*CourseInstance, error) {
	return &CourseInstance{
		InstanceID:          instanceId,
		Year:                year,
		CourseID:            courseId,
		SemesterID:          semesterId,
		AcademicYearID:      academicYearId,
		HardnessCoefficient: hardnessCoefficient,
		Form:                form,
		Mode:                mode,
		GroupsNeeded:        groupsNeeded,
		GroupsTaken:         groupsTaken,
		PIAllocationStatus:  piAllocStat,
		TIAllocationStatus:  tiAllocStatus,
	}, nil
}
