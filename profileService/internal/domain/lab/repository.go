package lab

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetAllLabs(ctx context.Context) ([]uuid.UUID, error)
	GetLabsByInstituteID(ctx context.Context, instituteID uuid.UUID) ([]uuid.UUID, error)
	GetLabByID(ctx context.Context, labID uuid.UUID) (*Lab, error)
}
