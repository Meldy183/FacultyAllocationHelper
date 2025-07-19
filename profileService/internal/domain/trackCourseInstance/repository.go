package trackcourseinstance

import "context"

type Repository interface {
	GetTracksOfCourseByInstanceID(ctx context.Context, instanceID int) ([]*TrackCourseInstance, error)
}
