package position

import "context"

type Repository interface {
	GetPositionByID(ctx context.Context, positionID int64) (*string, error)
	GetPositionIDByName(ctx context.Context, Name string) (*int64, error)
	GetAllPositions(ctx context.Context) ([]int64, error)
}
