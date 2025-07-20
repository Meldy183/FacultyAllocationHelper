package trackcourseinstance

import "context"

type Repository interface {
	GetTracksIDsOfCourseByInstanceID(ctx context.Context, instanceID int) ([]*TrackCourseInstance, error)
	AddTracksToCourseInstance(ctx context.Context, instanceID int, trackIDs int) error
}
