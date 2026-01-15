package allocation

import (
	"context"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/staff"
)

type Service interface {
	AllocateFaculty(ctx context.Context, courseID int64, profileID int64, positionType *string, groupsAssigned *int64) (*staff.Staff, error)
	DeallocateFaculty(ctx context.Context, courseID int64, profileID int64, positionType *string, groupsAssigned *int64) error
}
