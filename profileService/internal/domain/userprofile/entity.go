package userprofile

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
}

func NewUserProfile(
	id int64,
	userID string,
	engName, russianName, alias string,
	startDate, endDate *time.Time,
) (*UserProfile, error) {
	return &UserProfile{
		ProfileID:   id,
		EnglishName: engName,
		RussianName: &russianName,
		Alias:       alias,
		StartDate:   startDate,
		EndDate:     endDate,
	}, nil
}
