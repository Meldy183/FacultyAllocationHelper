package position

import "context"

type Repository interface {
	GetPositionByID(ctx context.Context, positionID int) (*Position, error)
	GetAllPositions(ctx context.Context) ([]int64, error)
}
