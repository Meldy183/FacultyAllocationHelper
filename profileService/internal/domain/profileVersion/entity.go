package profileVersion

type ProfileVersion struct {
	ProfileVersionId int64
	ProfileID        int64
	Year             int
	Workload         *float64
	MaxLoad          *int
	PositionID       int
	EmploymentType   *string
	Degree           *bool
	Mode             *Mode
}

type Mode string
