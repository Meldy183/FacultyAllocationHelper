package courseInstitute

import "context"

type Repository interface {
	GetCourseInstituteByID(ctx context.Context, instituteID int64) (*CourseInstitute, error)
	GetAllCourseInstitutes(ctx context.Context) ([]*CourseInstitute, error)
}
