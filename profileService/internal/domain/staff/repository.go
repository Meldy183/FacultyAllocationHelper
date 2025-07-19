package staff

import "context"

type Repository interface {
	GetAllStaffByInstanceID(ctx context.Context, instanceID int64) ([]*Staff, error)
	AddStaff(ctx context.Context, staff *Staff) error
	UpdateStaff(ctx context.Context, staff *Staff) error
}
