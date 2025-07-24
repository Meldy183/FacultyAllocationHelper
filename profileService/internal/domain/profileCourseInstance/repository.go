package profileCourseInstance

import (
	"context"
)

type Repository interface {
	GetCourseInstancesByVersionID(ctx context.Context, versionID int64) ([]int64, error)
	AddCourseInstance(ctx context.Context, profileVersionToCourseInstance *ProfileVersionCourseInstance) error
}
