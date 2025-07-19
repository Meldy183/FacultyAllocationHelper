package staff

import "context"

type Service interface {
	GetAllStaffByInstanceID(ctx context.Context, instanceID int64) ([]*Staff, error)
	AddStaff(ctx context.Context, staff *Staff) error
}
