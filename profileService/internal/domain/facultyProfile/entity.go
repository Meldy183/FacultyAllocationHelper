package facultyProfile

import (
	"time"
)

type UserProfile struct {
	ProfileID   int64
	Email       string
	EnglishName string
	RussianName *string
	Alias       string
	StartDate   *time.Time
	EndDate     *time.Time
	Fsro        *string
	Status      *string
}

func NewUserProfile(
	profileID int64,
	engName, russianName, alias string,
	startDate, endDate *time.Time,
) *UserProfile {
	return &UserProfile{
		ProfileID:   profileID,
		EnglishName: engName,
		RussianName: &russianName,
		Alias:       alias,
		StartDate:   startDate,
		EndDate:     endDate,
	}
}
