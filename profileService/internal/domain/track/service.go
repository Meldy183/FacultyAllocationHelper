package track

import (
	"context"
)

type Service interface {
	GetAllTracks(ctx context.Context) ([]*Track, error)
	GetTrackNameByID(ctx context.Context, trackID int64) (*string, error)
	GetTracksOfCourseByInstanceID(ctx context.Context, instanceID int64) ([]int64, error)
	GetTracksNamesOfCourseByCourseInstanceID(ctx context.Context, instanceID int64) ([]*string, error)
}
