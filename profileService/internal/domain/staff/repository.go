package staff

import "context"

type Repository interface {
	GetStaffByInstanceID(ctx context.Context, instanceID int) ([]*Staff, error)
	AddStaff(ctx context.Context, staff *Staff) error
	UpdateStaff(ctx context.Context, staff *Staff) error
}
