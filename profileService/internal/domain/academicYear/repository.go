package academicYear

import "context"

type Repository interface {
	GetAcademicYearNameByID(ctx context.Context, yearID *int64) (*string, error)
}
