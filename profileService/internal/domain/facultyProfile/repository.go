package facultyProfile

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	AddProfile(ctx context.Context, profile *UserProfile) error
	GetProfileByID(ctx context.Context, profileID uuid.UUID) (*UserProfile, error)
	UpdateProfileByID(ctx context.Context, profile *UserProfile) error
	GetProfileIDsByInstituteIDs(ctx context.Context, instituteIDs []uuid.UUID) ([]uuid.UUID, error)
	GetProfileIDsByPositionIDs(ctx context.Context, positionIDs []uuid.UUID) ([]uuid.UUID, error)
}
