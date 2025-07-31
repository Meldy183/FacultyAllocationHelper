package profileVersion

type ProfileVersion struct {
	ProfileVersionId int64
	ProfileID        int64
	Year             int64
	MaxLoad          *int64
	PositionID       int64
	EmploymentType   *string
	StudentType      *string
	Fsro             *string
	Degree           *bool
	Mode             *string
	FrontalHours     *int64
	ExtraActivities  *float64
}
