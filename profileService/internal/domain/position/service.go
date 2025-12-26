package position

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetPositionByID(ctx context.Context, positionID uuid.UUID) (*string, error)
	GetPositionIDByName(ctx context.Context, Name string) (*uuid.UUID, error)
	GetAllPositions(ctx context.Context) ([]uuid.UUID, error)
}
