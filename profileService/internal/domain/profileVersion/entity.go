package profileVersion

import "github.com/google/uuid"

type ProfileVersion struct {
	ProfileVersionId uuid.UUID
	ProfileID        uuid.UUID
	Year             int64
	MaxLoad          *int64
	PositionID       uuid.UUID
	EmploymentType   *string
	StudentType      *string
	Fsro             *string
	Degree           *bool
	Mode             *string
	FrontalHours     *int64
	ExtraActivities  *float64
}
