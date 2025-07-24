package staff

type Staff struct {
	AssignmentID     int64
	InstanceID       int64
	ProfileVersionID int64
	PositionType     *string
	GroupsAssigned   *int64
	IsConfirmed      bool
	LecturesCount    *int64
	TutorialsCount   *int64
	LabsCount        *int64
}
