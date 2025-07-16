package courseStaff

type UserProfile struct {
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

func NewUserProfile(
	AssignmentID int,
	InstanceID int,
	ProfileVersionID int,
	PositionType *string,
	ContributionCoefficient *int,
	GroupsAssigned *int,
	IsConfirmed bool,
	LabsCount *int,
	TutorialsCount *int,
	LecturesCount *int,
) *UserProfile {
	return &UserProfile{
		AssignmentID:            AssignmentID,
		InstanceID:              InstanceID,
		ProfileVersionID:        ProfileVersionID,
		PositionType:            PositionType,
		ContributionCoefficient: ContributionCoefficient,
		GroupsAssigned:          GroupsAssigned,
		IsConfirmed:             IsConfirmed,
		LabsCount:               LabsCount,
		TutorialsCount:          TutorialsCount,
		LecturesCount:           LecturesCount,
	}
}
