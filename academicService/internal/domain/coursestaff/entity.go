package coursestaff

import "github.com/shopspring/decimal"

type CourseStaff struct {
	AssignmentID            int64
	InstanceID              int64
	ProfileID               int64
	PositionType            string
	ContributionCoefficient decimal.Decimal
	GroupsAssigned          int8
	IsConfirmed             bool
	LabsCount               int8
	TutorialsCount          int8
	LecturesCount           int8
}
