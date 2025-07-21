package staff

type Staff struct {
	AssignmentID     int
	InstanceID       int64
	ProfileVersionID int64
	PositionType     *string
	GroupsAssigned   *int
	IsConfirmed      bool
	LecturesCount    *int
	TutorialsCount   *int
	LabsCount        *int
}
