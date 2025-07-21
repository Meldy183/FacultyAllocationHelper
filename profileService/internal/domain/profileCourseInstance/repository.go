package profileCourseInstance

import "context"

type Repository interface {
	GetCourseInstancesByVersionID(ctx context.Context, profileID int64) ([]int64, error)
	AddCourseInstance(ctx context.Context, userCourseInstance *ProfileCourseInstance) error
}
