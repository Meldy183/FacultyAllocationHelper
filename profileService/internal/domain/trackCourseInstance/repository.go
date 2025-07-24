package trackcourseinstance

import "context"

type Repository interface {
	GetTracksIDsOfCourseByInstanceID(ctx context.Context, instanceID int64) ([]int64, error)
	AddTracksToCourseInstance(ctx context.Context, instanceID int64, trackIDs int64) error
}
