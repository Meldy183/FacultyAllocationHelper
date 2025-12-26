package profileVersion

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetVersionByProfileID(ctx context.Context, profileID uuid.UUID, year int64) (*ProfileVersion, error)
	GetVersionIDByProfileID(ctx context.Context, profileID uuid.UUID, year int64) (uuid.UUID, error)
	AddProfileVersion(ctx context.Context, version *ProfileVersion) error
	GetVersionByVersionID(ctx context.Context, versionID uuid.UUID) (*ProfileVersion, error)
}
