package trackcourseinstance

import "context"

type Service interface {
	GetTracksIDsOfCourseByInstanceID(ctx context.Context, instanceID int) ([]*TrackCourseInstance, error)
	AddTracksToCourseInstance(ctx context.Context, instanceID, trackID int) error
}
