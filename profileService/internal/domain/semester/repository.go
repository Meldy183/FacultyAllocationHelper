package semester

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetSemesterNameByID(ctx context.Context, semesterID uuid.UUID) (*string, error)
	GetAllSemesters(ctx context.Context) ([]Semester, error)
}
