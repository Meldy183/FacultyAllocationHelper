package trackcourseinstance

import "context"

type Service interface {
	GetTracksIDsOfCourseByInstanceID(ctx context.Context, instanceID int64) ([]int64, error)
	AddTracksToCourseInstance(ctx context.Context, instanceID, trackID int64) error
}
