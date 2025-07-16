package userprofileversion

// import "errors"

// var (
// 	ErrInvalidMode     = errors.New("userprofile: invalid work mode")
// 	ErrInvalidMaxload  = errors.New("userprofile: max load must be in [0,40]")
// 	ErrInvalidWorkload = errors.New("userprofile: workload must be in [0,1]")
// 	ErrInvalidPosition = errors.New("invalid position of a person")
// )

type UserProfileVersion struct {
	ProfileID      int64
	Year           int
	Semester       string
	LecturesCount  *int
	TutorialsCount *int
	LabsCount      *int
	ElectivesCount *int
	Workload       *float64
	Maxload        *int
	PositionID     int
	EmploymentType *string
	Degree         *bool
	Mode           *Mode
}

type Mode string

const (
	ModeOnsite Mode = "Onsite"
	ModeMixed  Mode = "Mixed"
	ModeRemote Mode = "Remote"
)

func NewUserProfileVersion(
	profileID int64,
	year, positionID int,
	semester string,
	lectures_count, tutorials_count, labs_count, electives_count, maxload *int,
	workload *float64,
	employmentType *string,
	degree *bool,
	mode *Mode,
) *UserProfileVersion {
	// if !IsMaxloadValid(*maxload) {
	// 	return nil, ErrInvalidWorkload
	// }
	// if !IsWorkloadValid(*workload) {
	// 	return nil, ErrInvalidWorkload
	// }
	// if !IsPositionValid(positionID) {
	// 	return nil, ErrInvalidPosition
	// }
	return &UserProfileVersion{
		ProfileID:      profileID,
		Year:           year,
		Semester:       semester,
		LecturesCount:  lectures_count,
		TutorialsCount: tutorials_count,
		LabsCount:      labs_count,
		ElectivesCount: electives_count,
		Workload:       workload,
		Maxload:        maxload,
		PositionID:     positionID,
		EmploymentType: employmentType,
		Degree:         degree,
		Mode:           mode,
	}
}

// func IsMaxloadValid(maxload int) bool {
// 	return maxload <= 40 && maxload >= 0
// }

// func IsWorkloadValid(workload float64) bool {
// 	return workload <= 1 && workload >= 0
// }

// func IsPositionValid(positionID int) bool {
// 	return positionID <= 7 && positionID >= 1
// }
