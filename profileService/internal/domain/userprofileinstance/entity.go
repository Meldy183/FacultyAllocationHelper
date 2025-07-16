package userprofileversion

import (
	"errors"
)

var (
	// ErrInvalidMode is returned when NewUserProfile receives an unsupported mode.
	ErrInvalidMode = errors.New("userprofile: invalid work mode")

	// ErrInvalidLoad is returned when MaxLoad is not from 0 to 40.
	ErrInvalidMaxload  = errors.New("userprofile: max load must be in [0,40]")
	ErrInvalidWorkload = errors.New("userprofile: workload must be in [0,1]")
	ErrInvalidPosition = errors.New("invalid position of a person")
)

type UserProfileVersion struct {
	ProfileID       int64
	Year            int
	Semester        string
	Lectures_count  *int
	Tutorials_count *int
	Labs_count      *int
	Electives_count *int
	Workload        *float64
	Maxload         *int
	PositionID      int
	EmploymentType  *string
	Degree          *bool
	Mode            *Mode
}

type Mode string

const (
	ModeOnsite Mode = "onsite"
	ModeMixed  Mode = "mixed"
	ModeRemote Mode = "remote"
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
) (*UserProfileVersion, error) {
	if !IsMaxloadValid(*maxload) {
		return nil, ErrInvalidWorkload
	}
	if !IsWorkloadValid(*workload) {
		return nil, ErrInvalidWorkload
	}
	if !IsPositionValid(positionID) {
		return nil, ErrInvalidPosition
	}
	return &UserProfileVersion{
		ProfileID:       profileID,
		Year:            year,
		Semester:        semester,
		Lectures_count:  lectures_count,
		Tutorials_count: tutorials_count,
		Labs_count:      labs_count,
		Electives_count: electives_count,
		Workload:        workload,
		Maxload:         maxload,
		PositionID:      positionID,
		EmploymentType:  employmentType,
		Degree:          degree,
		Mode:            mode,
	}, nil
}

func IsMaxloadValid(maxload int) bool {
	return maxload <= 40 && maxload >= 0
}

func IsWorkloadValid(workload float64) bool {
	return workload <= 1 && workload >= 0
}

func IsPositionValid(positionID int) bool {
	return positionID <= 7 && positionID >= 1
}
