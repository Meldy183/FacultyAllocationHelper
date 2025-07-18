package track

import "context"

type Repository interface {
	GetAllTracks(ctx context.Context) (*[]Track, error)
	GetTrackNameByID(ctx context.Context, trackID int64) (*string, error)
}
