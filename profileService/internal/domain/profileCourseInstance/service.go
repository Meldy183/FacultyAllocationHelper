package profileCourseInstance

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetCourseInstancesByVersionID(ctx context.Context, profileID uuid.UUID) ([]uuid.UUID, error)
	AddCourseInstance(ctx context.Context, userCourseInstance *ProfileVersionCourseInstance) error
}
