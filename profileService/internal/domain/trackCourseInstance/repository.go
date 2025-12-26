package trackcourseinstance

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetTracksIDsOfCourseByInstanceID(ctx context.Context, instanceID uuid.UUID) ([]uuid.UUID, error)
	AddTracksToCourseInstance(ctx context.Context, instanceID uuid.UUID, trackIDs uuid.UUID) error
}
