package userprofile

import (
	"errors"
	"time"
)

var (
	// ErrInvalidMode is returned when NewUserProfile receives an unsupported mode.
	ErrInvalidMode = errors.New("userprofile: invalid work mode")

	// ErrInvalidLoad is returned when MaxLoad is not from 0 to 40.
	ErrInvalidLoad = errors.New("userprofile: max load must be in [0,40]")

	ErrInvalidPosition = errors.New("invalid position of a person")
)

type UserProfile struct {
	ProfileID      int64
	UserID         string
	Email          string
	Workload       float64
	Position       Position
	EnglishName    string
	RussianName    string
	Alias          string
	EmploymentType string
	StudentType    string
	Degree         bool
	Mode           Mode
	StartDate      *time.Time
	EndDate        *time.Time
	MaxLoad        int
}

type Mode string
type Position string

const (
	ModeOnsite Mode = "onsite"
	ModeMixed  Mode = "mixed"
	ModeRemote Mode = "remote"
)

const (
	PositionTA         Position = "TA"
	PositionIntern     Position = "TA Intern"
	PositionInstructor Position = "Instructor"
	PositionSenior     Position = "Senior Instructor"
	PositionDocent     Position = "Docent"
	PositionProfessor  Position = "Professor"
	PositionVisiting   Position = "Visiting"
)

func NewUserProfile(
	id int64,
	userID string,
	position Position,
	engName, russianName, alias, employmentType string,
	degree bool,
	mode Mode,
	startDate, endDate *time.Time,
	maxLoad int,
) (*UserProfile, error) {
	if !mode.IsValid() {
		return nil, ErrInvalidMode
	}
	if !IsWorkloadValid(maxLoad) {
		return nil, ErrInvalidLoad
	}
	if !IsPositionValid(position) {
		return nil, ErrInvalidPosition
	}
	return &UserProfile{
		ProfileID:      id,
		UserID:         userID,
		Position:       position,
		EnglishName:    engName,
		RussianName:    russianName,
		Alias:          alias,
		EmploymentType: employmentType,
		Degree:         degree,
		Mode:           mode,
		StartDate:      startDate,
		EndDate:        endDate,
		MaxLoad:        maxLoad,
	}, nil
}

func IsWorkloadValid(workload int) bool {
	return workload <= 40 && workload >= 0
}

func IsPositionValid(pos Position) bool {
	switch pos {
	case PositionInstructor, PositionSenior, PositionTA,
		PositionIntern, PositionDocent, PositionProfessor, PositionVisiting:
		return true
	default:
		return false
	}
}

func (m Mode) IsValid() bool {
	switch m {
	case ModeOnsite, ModeMixed, ModeRemote:
		return true
	default:
		return false
	}
}

func (p *UserProfile) ChangeMode(m Mode) error {
	if !m.IsValid() {
		return ErrInvalidMode
	}
	p.Mode = m
	return nil
}
