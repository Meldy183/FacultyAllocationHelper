package language

import "context"

type Repository interface {
	GetAllLanguages(ctx context.Context) ([]string, error)
	GetLanguageByCode(ctx context.Context, code string) (*Language, error)
	GetCodeByLanguageName(ctx context.Context, name string) (*string, error)
}
