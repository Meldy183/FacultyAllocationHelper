package userinstitute

import "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/institute"

type Repository interface {
	GetUserInstituteByID(ProfileID int64) (institute.Institute, error)
	AddUserInstitute(ProfileID int64, InstituteID int64, isReps bool) error
}
