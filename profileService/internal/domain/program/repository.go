package program

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetAllPrograms(ctx context.Context) ([]*Program, error)
	GetProgramNameByID(ctx context.Context, id uuid.UUID) (*string, error)
	GetProgramIDByName(ctx context.Context, name string) (*uuid.UUID, error)
}
