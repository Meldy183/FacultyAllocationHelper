package position

import "context"

type Repository interface {
	GetByID(ctx context.Context, positionID int) (*string, error)
	GetAll(ctx context.Context) ([]*Position, error)
}
