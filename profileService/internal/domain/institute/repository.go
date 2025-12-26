package institute

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetInstituteByID(ctx context.Context, instituteID uuid.UUID) (*Institute, error)
	GetAllInstitutes(ctx context.Context) ([]*Institute, error)
	GetInstituteIDByName(ctx context.Context, instituteName string) (*uuid.UUID, error)
}
