package academicYear

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetAcademicYearNameByID(ctx context.Context, yearID uuid.UUID) (*string, error)
	GetAllAcademicYears(ctx context.Context) ([]AcademicYear, error)
}
