package language

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]*Language, error)
	GetByCode(ctx context.Context, code string) (*Language, error)
}
