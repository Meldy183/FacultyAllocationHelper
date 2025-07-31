package profileLanguage

import (
	"context"
)

type Repository interface {
	AddUserLanguage(ctx context.Context, userLanguage *ProfileLanguage) error
	GetProfileLanguages(ctx context.Context, profileID int64) ([]string, error)
}
