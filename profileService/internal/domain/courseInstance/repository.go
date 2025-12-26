package courseInstance

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	AddNewCourseInstance(ctx context.Context, course *CourseInstance) error
	GetCourseInstanceByID(ctx context.Context, courseID uuid.UUID) (*CourseInstance, error)
	UpdateCourseInstanceByID(ctx context.Context, id uuid.UUID, course *CourseInstance) error
	GetInstancesIDsByInstituteIDs(ctx context.Context, instituteIDs []uuid.UUID) ([]uuid.UUID, error)
	GetInstancesIDsByAcademicYearIDs(ctx context.Context, academicYearIDs []uuid.UUID) ([]uuid.UUID, error)
	GetInstancesIDsBySemesterIDs(ctx context.Context, semesterIDs []uuid.UUID) ([]uuid.UUID, error)
	GetInstancesIDsByProgramIDs(ctx context.Context, programIDs []uuid.UUID) ([]uuid.UUID, error)
	GetInstancesByAllocationStatus(ctx context.Context) ([]uuid.UUID, error)
	GetInstancesByYear(ctx context.Context, year int64) ([]uuid.UUID, error)
	GetInstancesByVersionID(ctx context.Context, versionID uuid.UUID) ([]uuid.UUID, error)
	GetAllInstancesIDs(ctx context.Context) ([]uuid.UUID, error)
}
