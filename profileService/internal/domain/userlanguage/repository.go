package userlanguage

import (
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/language"
)

type Repository interface {
	Add(ProfileID int64, languageCode string) error
	GetUserLanguages(ProfileID int64) ([]*language.Language, error)
}
