package trackcourseinstance

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetTracksIDsOfCourseByInstanceID(ctx context.Context, instanceID uuid.UUID) ([]uuid.UUID, error)
	AddTracksToCourseInstance(ctx context.Context, instanceID, trackID uuid.UUID) error
}
