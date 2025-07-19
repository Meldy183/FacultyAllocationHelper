package track

import (
	"context"
	trackcourseinstance "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/trackCourseInstance"
)

type Service interface {
	GetTracksOfCourseByInstanceID(ctx context.Context, instanceID int) ([]*trackcourseinstance.TrackCourseInstance, error)
	GetAllTracks(ctx context.Context) (*[]Track, error)
	GetTrackNameByID(ctx context.Context, trackID int64) (*string, error)
	ConvertTrackCourseInstanceToTrackNames(context.Context, []*trackcourseinstance.TrackCourseInstance) ([]*string, error)
	GetTracksNamesOfCourseByCourseID(ctx context.Context, instanceID int) ([]*string, error)
}
