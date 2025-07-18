package staff

type Staff struct {
	AssignmentID            int
	InstanceID              int
	ProfileVersionID        int
	PositionType            *string
	ContributionCoefficient *int
	GroupsAssigned          *int
	IsConfirmed             bool
	LabsCount               *int
	TutorialsCount          *int
	LecturesCount           *int
}
