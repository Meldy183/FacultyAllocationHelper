package trackcourseinstance

import "context"

type Repository interface {
	GetTracksIDsOfCourseByInstanceID(ctx context.Context, instanceID int) ([]*TrackCourseInstance, error)
}
