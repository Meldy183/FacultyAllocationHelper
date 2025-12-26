package staff

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetAllStaffByInstanceID(ctx context.Context, instanceID uuid.UUID) ([]*Staff, error)
	AddStaff(ctx context.Context, staff *Staff) error
	GetPI(staff []*Staff) *Staff
	GetTI(staff []*Staff) *Staff
	GetTAs(staff []*Staff) []*Staff
}
