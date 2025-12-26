package courseInstance

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetCourseInstanceByID(ctx context.Context, courseID uuid.UUID) (*CourseInstance, error)
	AddCourseInstance(ctx context.Context, course *CourseInstance) error
	UpdateCourseInstanceByID(ctx context.Context, id uuid.UUID, course *CourseInstance) error
	GetInstancesByInstituteIDs(ctx context.Context, instituteIDs []uuid.UUID) ([]uuid.UUID, error)
	GetInstancesByAcademicYearIDs(ctx context.Context, academicYearIDs []uuid.UUID) ([]uuid.UUID, error)
	GetInstancesBySemesterIDs(ctx context.Context, semesterIDs []uuid.UUID) ([]uuid.UUID, error)
	GetInstancesByProgramIDs(ctx context.Context, programIDs []uuid.UUID) ([]uuid.UUID, error)
	GetInstancesByAllocationStatus(ctx context.Context, allocNotFinished bool) ([]uuid.UUID, error)
	GetInstancesByYear(ctx context.Context, year int64) ([]uuid.UUID, error)
	GetInstancesByVersionID(ctx context.Context, versionID uuid.UUID) ([]uuid.UUID, error)
}
