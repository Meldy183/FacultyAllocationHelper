package courseinstance

type CourseInstance struct {
	InstanceID         int64
	CourseID           int64
	Semester           string
	Year               int8
	LabFormat          string
	LectureFormat      string
	AcademicYear       string
	Form               string
	GroupsNeeded       int8
	GroupsTaken        int8
	PIAllocationStatus string
	TIAllocationStatus string
}
