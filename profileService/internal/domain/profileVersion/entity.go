package profileVersion

type ProfileVersion struct {
	ProfileID      int64
	Year           int
	Semester       string
	Workload       *float64
	MaxLoad        *int
	PositionID     int
	EmploymentType *string
	Degree         *bool
	Mode           *Mode
}

type Mode string

func NewUserProfileVersion(
	profileID int64,
	year, positionID int,
	semester string,
	maxLoad *int,
	workload *float64,
	employmentType *string,
	degree *bool,
	mode *Mode,
) *ProfileVersion {
	return &ProfileVersion{
		ProfileID:      profileID,
		Year:           year,
		Semester:       semester,
		Workload:       workload,
		MaxLoad:        maxLoad,
		PositionID:     positionID,
		EmploymentType: employmentType,
		Degree:         degree,
		Mode:           mode,
	}
}
