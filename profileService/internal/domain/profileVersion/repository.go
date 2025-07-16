package profileVersion

import "context"

type Repository interface {
	AddProfileVersion(ctx context.Context, profile *ProfileVersion) error
	GetProfileVersionByID(ctx context.Context, profileID int64) (*ProfileVersion, error)
	UpdateProfileVersionByID(ctx context.Context, profile *ProfileVersion) error
	GetProfileVersionsByFilter(ctx context.Context, institutes []int, positions []int) ([]int64, error)
}
