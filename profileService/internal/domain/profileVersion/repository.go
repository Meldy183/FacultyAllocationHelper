package profileVersion

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	AddProfileVersion(ctx context.Context, profile *ProfileVersion) error
	GetVersionByProfileID(ctx context.Context, profileID uuid.UUID, year int64) (*ProfileVersion, error)
	GetVersionByVersionID(ctx context.Context, profileID uuid.UUID) (*ProfileVersion, error)
}

//TODO: UpdateProfileVersionByID(ctx context.Context, profile *ProfileVersion) error
