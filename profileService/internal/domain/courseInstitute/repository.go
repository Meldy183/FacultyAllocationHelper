package courseInstitute

import "context"

type Repository interface {
	GetCourseInstituteByID(ctx context.Context, instituteID int64) (*InstituteCourseLink, error)
	GetAllCourseInstitutes(ctx context.Context) ([]*InstituteCourseLink, error)
}
