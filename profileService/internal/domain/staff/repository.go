package staff

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetAllStaffByInstanceID(ctx context.Context, instanceID uuid.UUID) ([]*Staff, error)
	AddStaff(ctx context.Context, staff *Staff) error
	UpdateStaff(ctx context.Context, staff *Staff) error
}
