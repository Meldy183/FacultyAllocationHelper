package profileCourseInstance

import "context"

type Service interface {
	GetCourseInstancesByVersionID(ctx context.Context, profileID int64) ([]int64, error)
	AddCourseInstance(ctx context.Context, userCourseInstance *ProfileCourseInstance) error
}
