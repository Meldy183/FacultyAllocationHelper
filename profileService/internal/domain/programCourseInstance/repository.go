package programcourseinstance

import "context"

type Repository interface {
	GetPorgramCourseInstancesByID(ctx context.Context, code string) (*PorgramCourseInstance, error)
}
