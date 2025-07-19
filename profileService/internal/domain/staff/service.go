package staff

import "context"

type Service interface {
	GetAllStaffByInstanceID(ctx context.Context, instanceID int) ([]*Staff, error)
	AddStaff(ctx context.Context, staff *Staff) error
}
