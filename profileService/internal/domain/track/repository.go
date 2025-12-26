package track

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetAllTracks(ctx context.Context) ([]*Track, error)
	GetTrackNameByID(ctx context.Context, trackID uuid.UUID) (*string, error)
	GetTrackIDByName(ctx context.Context, trackName string) (*uuid.UUID, error)
}
