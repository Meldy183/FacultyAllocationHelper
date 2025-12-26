package track

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetAllTracks(ctx context.Context) ([]*Track, error)
	GetTrackNameByID(ctx context.Context, trackID uuid.UUID) (*string, error)
	GetTrackIDByName(ctx context.Context, trackName string) (*uuid.UUID, error)
	GetTracksOfCourseByInstanceID(ctx context.Context, instanceID uuid.UUID) ([]uuid.UUID, error)
	GetTracksNamesOfCourseByCourseInstanceID(ctx context.Context, instanceID uuid.UUID) ([]*string, error)
}
