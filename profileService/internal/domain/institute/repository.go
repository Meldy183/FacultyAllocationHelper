package institute

import "context"

type Repository interface {
	GetByID(ctx context.Context, instituteID int64) (*Institute, error)
	GetAll(ctx context.Context) ([]*Institute, error)
}
