package profileVersion

type ProfileVersion struct {
	ProfileVersionId int64
	ProfileID        int64
	Year             int
	MaxLoad          *int
	PositionID       int
	EmploymentType   *string
	StudentType      *string
	Fsro             *string
	Degree           *bool
	Mode             *string
	FrontalHours     *int
	ExtraActivities  *float64
}
