package profileLanguage

import "github.com/google/uuid"

type ProfileLanguage struct {
	ProfileLanguageID uuid.UUID
	ProfileID         uuid.UUID
	LanguageCode      string
}
