package profileVersion

import "context"

type Repository interface {
	AddProfileVersion(ctx context.Context, profile *ProfileVersion) error
	GetVersionByProfileID(ctx context.Context, profileID int64, year int) (*ProfileVersion, error)
	GetVersionByVersionID(ctx context.Context, profileID int64) (*ProfileVersion, error)
}

//TODO: UpdateProfileVersionByID(ctx context.Context, profile *ProfileVersion) error
