package facultyProfile

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	AddProfile(ctx context.Context, profile *UserProfile) error
	GetProfileByID(ctx context.Context, profileID uuid.UUID) (*UserProfile, error)
	UpdateProfileByID(ctx context.Context, profile *UserProfile) error
	GetProfilesByFilters(ctx context.Context, institutes []uuid.UUID, positions []uuid.UUID) ([]uuid.UUID, error)
}
