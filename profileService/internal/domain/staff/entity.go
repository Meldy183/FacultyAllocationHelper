package staff

type Staff struct {
	AssignmentID            int
	InstanceID              int
	ProfileVersionID        int
	PositionType            *string
	ContributionCoefficient *int
	GroupsAssigned          *int
	IsConfirmed             bool
	LecturesCount           *int
	TutorialsCount          *int
	LabsCount               *int
}
