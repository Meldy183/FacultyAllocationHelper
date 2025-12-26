package profileLanguage

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	AddUserLanguage(ctx context.Context, userLanguage *ProfileLanguage) error
	GetProfileLanguages(ctx context.Context, profileID uuid.UUID) ([]string, error)
}
