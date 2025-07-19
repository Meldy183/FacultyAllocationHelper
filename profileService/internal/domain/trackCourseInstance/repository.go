package trackcourseinstance

import "context"

type Repository interface {
	GetTracksOfCourseByCourseID(ctx context.Context, courseID int) ([]*TrackCourseInstance, error)
}
