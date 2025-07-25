package completeuser

import (
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
)

type CompleteUser struct {
	Profile facultyProfile.UserProfile
	Version profileVersion.ProfileVersion
}
