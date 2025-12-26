package profileCourseInstance

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetCourseInstancesByVersionID(ctx context.Context, versionID uuid.UUID) ([]uuid.UUID, error)
	AddCourseInstance(ctx context.Context, profileVersionToCourseInstance *ProfileVersionCourseInstance) error
}
