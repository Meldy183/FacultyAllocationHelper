package position

import "context"

type Service interface {
	GetPositionByID(ctx context.Context, positionID int) (*string, error)
	GetAllPositions(ctx context.Context) ([]*Position, error)
}
