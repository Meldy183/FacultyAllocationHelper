package profileLanguage

import (
	"context"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/language"
)

type Service interface {
	AddUserLanguage(ctx context.Context, userLanguage *UserLanguage) error
	GetUserLanguages(ctx context.Context, profileID int64) ([]*language.Language, error)
}
