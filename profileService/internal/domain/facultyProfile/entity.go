package facultyProfile

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	ProfileID   uuid.UUID
	Email       string
	EnglishName string
	RussianName *string
	Alias       string
	StartDate   *time.Time
	EndDate     *time.Time
	Status      *string
}
