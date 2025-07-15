package position

import "context"

type Repository interface {
	GetPositionByID(ctx context.Context, positionID int) (*string, error)
	GetAllPositions(ctx context.Context) ([]*Position, error)
}
