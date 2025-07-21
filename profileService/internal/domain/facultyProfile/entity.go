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
