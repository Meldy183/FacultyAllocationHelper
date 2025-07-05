package usercourseinstance

import "context"

type Repository interface {
	GetInstancesByProfileID(ctx context.Context, profileID int64) ([]int64, error)
	AddCourseInstance(ctx context.Context, profileID int64, instanceID int64) error
}
