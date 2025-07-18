package trackcourseinstance

import "context"

type Repository interface {
	GetTrackCourseInstancesByCourseID(ctx context.Context, courseID int) (*TrackCourseInstance, error)
}
