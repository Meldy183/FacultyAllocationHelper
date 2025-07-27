package CompleteUser

import (
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
)

type FullUser struct {
	UserProfile        facultyProfile.UserProfile
	UserProfileVersion profileVersion.ProfileVersion
	Institutes         []*string
	Languages          []*string
}
