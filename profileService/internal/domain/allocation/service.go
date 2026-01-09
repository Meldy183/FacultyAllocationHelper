package allocation

import "context"

type Service interface {
	AllocateFaculty(ctx context.Context, courseInstanceID int64, profileID int64, positionType string) error
	DeallocateFaculty(ctx context.Context, courseInstanceID int64, profileID int64, positionType string) error
}
