package staff

import "github.com/google/uuid"

type Staff struct {
	AssignmentID     uuid.UUID
	InstanceID       uuid.UUID
	ProfileVersionID uuid.UUID
	PositionType     *string
	GroupsAssigned   *int64
	IsConfirmed      bool
	LecturesCount    *int64
	TutorialsCount   *int64
	LabsCount        *int64
}
