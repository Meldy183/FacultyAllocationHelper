package program

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetProgramNamesByInstanceID(ctx context.Context, courseID uuid.UUID) ([]*string, error)
	GetProgramNameByID(ctx context.Context, id uuid.UUID) (*string, error)
	GetProgramIDByName(ctx context.Context, name string) (*uuid.UUID, error)
	GetAllPrograms(ctx context.Context) ([]*Program, error)
}
