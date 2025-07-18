package profileVersion

import (
	"context"
)

type Service interface {
	GetVersionByProfileID(ctx context.Context, profileID int64, year int) (*ProfileVersion, error)
	GetVersionIDByProfileID(ctx context.Context, profileID int64, year int) (int64, error)
	AddProfileVersion(ctx context.Context, version *ProfileVersion) error
}
