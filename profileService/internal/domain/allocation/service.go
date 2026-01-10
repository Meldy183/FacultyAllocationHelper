package allocation

import "context"

type Service interface {
	AllocateFaculty(ctx context.Context, courseInstanceID int64, profileID int64, positionType *string, groupsAssigned *int64) error
	DeallocateFaculty(ctx context.Context, courseInstanceID int64, profileID int64, positionType *string) error
}
